package service

type Service struct {
	key    []byte
	//bl     *bl.BusLogic
	health bool
	conSem chan bool
}

type Response struct {
	ReturnCode    int           `json:"code"`
	ReturnMessage string        `json:"message"`
	Data          interface{}   `json:"data"`
	//DevMessage    string        `json:"dev_message"`
	//context       []interface{} `json:"-"`
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

func (s *Service) NewErrorResponse(err error) (*Response, error) {
	if err == nil {
		panic("should not call this method with nil err")
	}
	derr, ok := err.(*de.QError)
	if ok {
		return &Response{
			ReturnCode:    derr.Code(),
			ReturnMessage: derr.DisplayMessage(),
			DevMessage:    derr.Error(),
		}, nil
	} else {
		return nil, err
	}
}