package gooo

import (
	"net/http"
)

type Cookies map[string]string

func NewCookies(c []*http.Cookie) *Cookies {
	r := make(Cookies, len(c))
	for _, v := range c {
		if v != nil {
			r[v.Name] = v.Value
		}
	}
	return &r
}
