package model

type Response struct {
	Error *ErrInfo    `json:"error"`
	Body  interface{} `json:"body"`
}
