package models

type Response struct {
	Status  string      `json:"status"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func GetResponse(status string, data interface{}, message string) Response {
	return Response{
		Status:  status,
		Data:    data,
		Message: message,
	}
}
