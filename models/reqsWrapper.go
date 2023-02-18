package models

type ReqsHeader struct {
	Name  string
	Value string
}

type ReqsCookie struct {
	Name  string
	Value string
}

type ReqsWrapper struct {
	ID          int
	Name        string
	Method      string
	URI         string
	Headers     []ReqsHeader
	Cookies     []ReqsCookie
	PayloadType string
	Payload     interface{}
}
