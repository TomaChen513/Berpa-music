package common

func SuccessMessage(message string, data interface{}) *ResponseBody {
	return &ResponseBody{
		Code:    200,
		Message: message,
		Success: true,
		Type:    "success",
		Data:    data,
	}
}
