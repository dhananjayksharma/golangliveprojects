package responses

type Response struct {
	Data      interface{}
	RecordSet interface{}
	Message   string
}

type ErrorMessage struct {
	Name  string `json:"name"`
	Desc  string `json:"desc"`
	Error string `json:"error"`
}
