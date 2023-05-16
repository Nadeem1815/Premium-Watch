package response

type Response struct{
	StatusCode int  `json:"status_code"`
	Message    string  `json:"message"`
	Data       interface{}`jsno:"data,omitempy"`
	Errors     interface{} `json:"errors,omitempty"`
}