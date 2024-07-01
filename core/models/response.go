package models

type Response struct {
	Status  string      `json:"status"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func GetResponse(status string, data interface{}, messages ...string) Response {
	var message string
	for _, m := range messages {
		message += " " + m
	}

	return Response{
		Status:  status,
		Data:    data,
		Message: message,
	}
}
