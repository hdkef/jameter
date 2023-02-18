package models

type ReqsWrapper struct {
	ID          int
	Name        string
	Method      string
	Headers     []string
	Cookies     []string
	PayloadType string
	Payload     interface{}
}
