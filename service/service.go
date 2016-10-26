package service

import (
	"audit-rule/buslogic"
	"github.com/kataras/iris"
	"encoding/json"
	"audit-rule/log"
	ae "audit-rule/error"
	"gopkg.in/validator.v2"
)

type Service struct {
	bl         *buslogic.BusLogic
}

func NewService() *Service {
	bl, err := buslogic.NewBusLogic()
	if err != nil {
		panic(err)
	}
	return &Service{
		bl:         bl,
	}
}

func (s *Service) RegisterAll(api *iris.Framework) {
	s.Register(api, "POST", "/antiFraud/buyer", s.test)
}

func (s *Service) Register(api *iris.Framework, method, path string, handler func(c *iris.Context) (*Response, error)) {
	log.CustomLogger.Debug("registered api: %s %s", method, path)
	api.HandleFunc(method, path, s.wrapIris(handler))
}

func (s *Service) wrapIris(callback func(c *iris.Context) (*Response, error)) iris.HandlerFunc {
	return func(c *iris.Context) {
		defer func() {
			err := recover()
			if err != nil {
				log.CustomLogger.Info("got panic for request: %s, err=%+v", string(c.RequestURI()), err)
				c.Write(`{
					"code": -1,
					"message": "系统错误",
					"dev_message": ""
				}`)
			}
		}()
		r, err := callback(c)
		if err != nil {
			log.CustomLogger.Error("got error err = %s", err.Error())
		}
		body, _ := json.Marshal(r)
		c.Write(string(body))
	}
}

func (s *Service) jsonRequest(c *iris.Context, request interface{}, callback func() (*Response, error)) (*Response, error) {
	err := json.Unmarshal(c.PostBody(), request)
	if err != nil {
		return s.NewResponse(ae.ERROR_REQUEST_INVALID, "error parse json from request", err.Error()), nil
	}
	errs := validator.Validate(request)
	if errs != nil {
		log.CustomLogger.ErrorWithoutStack("param missing: %v",errs)
		return s.NewResponse(ae.ERROR_MISSING_PARAM, "required param missing"), nil
	}
	return callback()
}
