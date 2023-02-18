package models

import "net/http"

type ResultChan struct {
	Res *http.Response
	Req *ReqsWrapper
}
