package main

import (
	"github.com/codegangsta/martini"
	"github.com/martini-contrib/render"
	"gooo"
	//"html/template"
	"net/http"
	"service/web/rendertext"
)

var BasePath = "../web"

//var HelperFuncs = template.FuncMap{
//	"text": func(s template.HTML) template.JS {
//		return template.JS(s)
//	},
//}

func Run(path string) {

	m := martini.Classic()
	m.Use(martini.Static(basepath("static")))

	m.Get("/:name.html", render.Renderer(render.Options{
		Directory: basepath("view/html"), // Specify what path to load the templates from.
		//Layout:          "layout",                       // Specify a layout template. Layouts can call {{ yield }} to render the current template.
		Extensions: []string{".tmpl", ".html"}, // Specify extensions to load for templates.
		//Funcs:           []template.FuncMap{HelperFuncs}, // Specify helper function maps for templates to access.
		Delims:          render.Delims{"{{", "}}"}, // Sets delimiters to the specified strings.
		Charset:         "UTF-8",                   // Sets encoding for json and html content-types. Default is "UTF-8".
		IndentJSON:      true,                      // Output human readable JSON
		IndentXML:       true,                      // Output human readable XML
		HTMLContentType: "application/xhtml+xml",   // Output XHTML content type instead of default "text/html"
	}), func(params martini.Params, r render.Render) {
		r.HTML(200, params["name"], 10, render.HTMLOptions{Layout: "layout"})
	})

	m.Get("/js/:name.js", rendertext.Renderer(rendertext.Options{
		Directory:  basepath("view/js"),
		Extensions: []string{".tmpl", ".js"},
		Delims:     rendertext.Delims{"{{", "}}"},
		Charset:    "UTF-8",
		IndentJSON: true,
		IndentXML:  true,
	}), func(params martini.Params, r rendertext.Render) {
		r.Text(200, params["name"], 10)
	})

	m.Get("/css/:name.js", rendertext.Renderer(rendertext.Options{
		Directory:  basepath("view/css"),
		Extensions: []string{".tmpl", ".css"},
		Delims:     rendertext.Delims{"{{", "}}"},
		Charset:    "UTF-8",
		IndentJSON: true,
		IndentXML:  true,
	}), func(params martini.Params, r rendertext.Render) {
		r.Text(200, params["name"], 10)
	})

	m.Get("/push", func(w http.ResponseWriter, r *http.Request) {
		c, ok := w.(http.Hijacker)
		if ok {
			conn, _, err := c.Hijack()
			if err != nil {
				return
			}
			conn.Write([]byte("hello push"))
		}
	})
	m.RunOnAddr(path)
}

func basepath(path ...string) string {
	return gooo.JoinPath(append([]string{BasePath}, path...)...)
}

func main() {
	Run(gooo.Port)
	return
}
