package route

import (
	"errors"
	"io/ioutil"
	"rego"
	"rego/server"
)

type CodeMap struct {
	Name   string
	Classs server.Classs
}

type CodeMaps []CodeMap

func NewCodeMaps() *CodeMaps {
	r := CodeMaps{}
	return &r
}

func (co *CodeMaps) Map(c1, c2, c3 byte) (m1, m2, m3 string, err error) {
	defer func() {
		if errr := recover(); errr != nil {
			err = errors.New("CodeMaps.Map: index out of range")
		}
	}()
	i1 := (*co)[c1]
	i2 := i1.Classs[c2]
	m3 = i2.Methods[c3]
	m2 = i2.Name
	m1 = i1.Name
	return
}

func (co *CodeMaps) Append(name string, classs server.Classs) {
	*co = append(*co, CodeMap{
		Name:   name,
		Classs: classs,
	})
}

func (co *CodeMaps) MakeReCodeMap() *ReCodeMaps {
	re := make(ReCodeMaps)
	for k1, v1 := range *co {
		for k2, v2 := range v1.Classs {
			for k3, v3 := range v2.Methods {
				re[v1.Name+"."+v2.Name+"."+v3] = uint32(k1|(k2<<8)|(k3<<16)) | (1 << 24)
			}
		}
	}
	return &re
}

func (co *CodeMaps) WriteFile(name string) {
	es := rego.NewEncodeStream()
	es.EnJson(co)
	ioutil.WriteFile(name, es.Bytes(), 0666)
}

func (co *CodeMaps) ReadFile(name string) {
	b, err := ioutil.ReadFile(name)
	if err != nil {
		rego.ERR(err)
		return
	}
	es := rego.NewEncodeStream()
	es.Set(b)
	es.DeJson(co)
}

type ReCodeMaps map[string]uint32

func (co *ReCodeMaps) Map(code string) (c1, c2, c3 byte, err error) {
	c := (*co)[code]
	if (uint32(c) >> 24) != 1 {
		return 0, 0, 0, errors.New(code + ": inexistence")
	}
	return byte(c), byte(c >> 8), byte(c >> 16), nil
}
