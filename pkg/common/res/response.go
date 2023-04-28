package res

type Response struct {
	StatusCode int         `json:"stastuscode"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	Errors     interface{} `json:"error"`
}
