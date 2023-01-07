package common

func AuthErrorMessage() *ResponseBody {
	return &ResponseBody{
		// Message: Message{
		Code:    304,
		Message: "static file path error",
		Success: false,
		Type:    "false",
		Data:    nil,
		// },
	}
}
