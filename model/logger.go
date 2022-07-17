package model

import (
	"net/http"
)

type Logger struct {
	Xid    string      `json:"xid"`
	Method string      `json:"method"`
	Url    string      `json:"url"`
	Header http.Header `json:"header"`
	Body   interface{} `json:"body"`
}
