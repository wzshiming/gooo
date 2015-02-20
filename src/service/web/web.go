package main

import (
	"github.com/codegangsta/martini"
	"github.com/martini-contrib/render"
	"gooo"
)

var BasePath = "../web"

func Run(path string) {
	m := martini.Classic()
	m.Use(martini.Static(basepath("static")))
	m.Use(render.Renderer(render.Options{
		Directory: basepath("view"), // Specify what path to load the templates from.
		//Layout:          "layout",                       // Specify a layout template. Layouts can call {{ yield }} to render the current template.
		Extensions: []string{".tmpl", ".html"}, // Specify extensions to load for templates.
		//Funcs:           []template.FuncMap{AppHelpers}, // Specify helper function maps for templates to access.
		Delims:          render.Delims{"{{", "}}"}, // Sets delimiters to the specified strings.
		Charset:         "UTF-8",                   // Sets encoding for json and html content-types. Default is "UTF-8".
		IndentJSON:      true,                      // Output human readable JSON
		IndentXML:       true,                      // Output human readable XML
		HTMLContentType: "application/xhtml+xml",   // Output XHTML content type instead of default "text/html"
	}))

	m.Get("/:name.html", func(params martini.Params, r render.Render) {
		r.HTML(200, params["name"], 10, render.HTMLOptions{Layout: "layout"})
	})

	m.Get("/js/:name.js", func(params martini.Params, r render.Render) {
		r.HTML(200, "js/"+params["name"], 10, render.HTMLOptions{Layout: "js/layout"})
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
