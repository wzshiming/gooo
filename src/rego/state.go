package rego

import ()

type LS uint32

const (
	LS_UNALLOW LS = iota
	LS_GUEST
	LS_LOGIN
	LS_PLAYING
	LS_INTERRU
	LS_OFFINE
)

type LF uint32

const (
	LF_COMM LF = 1 << (32 - 1 - iota)
	LF_VIP
	LF_ADMIN
	LF_NONE LF = 0
)

type Data struct {
	EncodeBytes
}

func (da *Data) LS() LS {
	var t struct {
		Ls___ LS
	}
	da.DeGob(&t)
	return t.Ls___
}
