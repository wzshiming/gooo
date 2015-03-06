package agent

import (
	"rego"
	//"net/http"
)

type Request struct {
	Request *rego.EncodeBytes
	Session *Session
}

type Response struct {
	Error    string
	Data     *rego.EncodeBytes
	Coverage *rego.EncodeBytes
	Response *rego.EncodeBytes
}
