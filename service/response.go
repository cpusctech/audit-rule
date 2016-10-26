package service

import (
	ae "audit-rule/error"
)

type Response struct {
	ReturnCode    int           `json:"code"`
	ReturnMessage string        `json:"message"`
	Data          interface{}   `json:"data"`
	DevMessage    string        `json:"dev_message"`
	context       []interface{} `json:"-"`
	Status        bool          `json:"status"`
	//Msg           string        `json:"msg"`
}

func (s *Service) NewApiOkResponse(extraData interface{}) *Response {
	return &Response{
		ReturnCode:    0,
		ReturnMessage: "OK",
		Data:          extraData,
	}
}

func (s *Service) NewAjaxOkResponse() *Response {
	return &Response{
		Status:    true,
		ReturnMessage: "OK",
	}
}

func (s *Service) NewAjaxErrorResponse(errorMessage string) *Response {
	return &Response{
		Status: false,
		ReturnMessage:    errorMessage,
	}
}

func (s *Service) NewResponse(code int, devMessage string, context ...interface{}) *Response {
	er := &Response{}
	er.ReturnCode = code
	er.ReturnMessage = ae.ReturnMessage(code, context...)
	er.DevMessage = devMessage
	er.context = context
	return er
}
