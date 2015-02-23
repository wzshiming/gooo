package rendertext

import (
	"github.com/oxtoacart/bpool"
	"github.com/wzshiming/ffmt"
	"text/template"
)

type Template struct {
	t    *template.Template
	bufs *bpool.BufferPool
}

func Compile(dir string, fun []template.FuncMap, ext ...string) *Template {
	return &Template{
		compile(
			Options{
				Directory:  dir,
				Funcs:      fun,
				Extensions: append([]string{".tmpl"}, ext...),
				Delims:     Delims{"{{", "}}"},
				Charset:    "UTF-8",
			},
		),
		bpool.NewBufferPool(64),
	}
}

func (r *Template) Get(name string, binding interface{}) string {
	b := r.bufs.Get()
	defer func() {
		r.bufs.Put(b)
	}()
	err := r.t.ExecuteTemplate(b, name, binding)
	if err != nil {
		return ""
	}
	return ffmt.Strip(string(b.Bytes()))
}
