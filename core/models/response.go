package models

type Response struct {
	Status  string      `json:"status"`
	Data    interface{} `json:"data"`
	Message []string    `json:"messages"`
}

func GetResponse(status string, data interface{}, messages ...string) Response {
	return Response{
		Status:  status,
		Data:    data,
		Message: messages,
	}
}
