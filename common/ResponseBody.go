package common

type ResponseBody struct {
	// Message Message `json:"message"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Type    string      `json:"type"`
}

// type Message struct {

// }
