package mail

type mail struct {
	SenderEmail string     `json:"senderEmail"`
	Templates   []Template `json:"templates"`
}

// template struct for storing templates information
// use templates to send the demo emails // TODO: try out this new feature of storing the templates.
type Template struct {
	Name    string            `json:"name"`
	Content string            `json:"content"`
	Subject string            `json:"subject"`
	Params  map[string]string `json:"params"`
}

func New(senderEmail string) *mail {
	return &mail{SenderEmail: senderEmail}
}
