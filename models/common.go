package models

//Response default response
type Response struct {
	StatusCode int         `json:"statuscode"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}
