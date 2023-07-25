package api

type JsonResult struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Success bool        `json:"success"`
}

func Json(code int, message string, data interface{}, success bool) *JsonResult {
	return &JsonResult{
		Code:    code,
		Message: message,
		Data:    data,
		Success: success,
	}
}

func JsonData(data interface{}) *JsonResult {
	return &JsonResult{
		Code:    200,
		Data:    data,
		Success: true,
		Message: "请求成功",
	}
}

func JsonSuccess() *JsonResult {
	return &JsonResult{
		Code:    200,
		Data:    nil,
		Success: true,
		Message: "请求成功",
	}
}

func JsonError(err *CodeError) *JsonResult {
	return &JsonResult{
		Code:    err.Code,
		Message: err.Message,
		Data:    err.Data,
		Success: false,
	}
}

func (json *JsonResult) JsonWithMsg(message string) *JsonResult {
	json.Message = message
	return json
}

func (json *JsonResult) JsonWithCode(code int) *JsonResult {
	json.Code = code
	return json
}

func (json *JsonResult) JsonWithData(data interface{}) *JsonResult {
	json.Data = data
	return json
}
