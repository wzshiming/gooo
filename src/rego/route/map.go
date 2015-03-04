package route

import (
	"errors"
	"rego"
	"rego/server"
)

type CodeMap struct {
	Name   string
	Classs server.Classs
}

type CodeMaps []CodeMap

func (co CodeMaps) Map(c1, c2, c3 byte) (m1, m2, m3 string, err error) {
	defer func() {
		if err := recover(); err != nil {
			rego.PANIC(err)
			return
		}
	}()
	i1 := co[c1]
	i2 := i1.Classs[c2]
	m3 = i2.Methods[c3]
	m2 = i2.Name
	m1 = i1.Name
	return
}

func (co CodeMaps) MakeReCodeMap() (re ReCodeMap) {
	re = make(ReCodeMap)
	for k1, v1 := range co {
		for k2, v2 := range v1.Classs {
			for k3, v3 := range v2.Methods {
				re[v1.Name+v2.Name+v3] = uint32(k1 | (k2 << 8) | (k3 << 16))
			}
		}
	}
	return
}

type ReCodeMap map[string]uint32

func (co ReCodeMap) ReCodeMap(code string) (c1, c2, c3 byte, err error) {
	c := co[code]
	if c == 0 {
		return 0, 0, 0, errors.New(code + ": inexistence")
	}
	return byte(c), byte(c >> 8), byte(c >> 16), nil
}
