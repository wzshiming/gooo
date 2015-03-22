package main

import (
	"branch/rendertext"
	"fmt"
	"github.com/go-martini/martini"
	"github.com/gorilla/websocket"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	ht "html/template"
	"net/http"
	"rego"
	"rego/agent"
	"rego/agent/defaul"
	"rego/cfg"
	"rego/route"
	tt "text/template"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}
var BasePath = "../web"

var cod *route.CodeMaps = route.NewCodeMaps()
var rec *route.ReCodeMaps

func init() {
	cod.ReadFile(defaul.MapFile)
	rec = cod.MakeReCodeMap()
}

var helpfunc = map[string]interface{}{
	"key": func(key string) string {
		c1, c2, c3, err := rec.Map(key)
		if err != nil {
			return "255,255,255"
		}
		return fmt.Sprintf("%d,%d,%d", c1, c2, c3)
	},
}

func run(ag *agent.Agent) {

	m := martini.Classic()

	m.Use(martini.Static(basepath("/static")))
	store := sessions.NewCookieStore([]byte("gooo"))
	m.Use(sessions.Sessions("session", store))

	m.Get("/:name.html", render.Renderer(render.Options{
		Directory:  basepath("/view/html"),     // Specify what path to load the templates from.
		Layout:     "layout",                   // Specify a layout template. Layouts can call {{ yield }} to render the current template.
		Extensions: []string{".tmpl", ".html"}, // Specify extensions to load for templates.
		Funcs:      []ht.FuncMap{helpfunc},     // Specify helper function maps for templates to access.
		Delims:     render.Delims{"{{", "}}"},  // Sets delimiters to the specified strings.
		Charset:    "UTF-8",                    // Sets encoding for json and html content-types. Default is "UTF-8".
		IndentJSON: true,                       // Output human readable JSON
		IndentXML:  true,                       // Output human readable XML
		//HTMLContentType: "application/xhtml+xml",    // Output XHTML content type instead of default "text/html"
	}), func(params martini.Params, r render.Render) {
		if params["name"] == "layout" || params["name"] == "" {
			params["name"] = "index"
		}
		r.HTML(200, params["name"], map[string]interface{}{}) //, render.HTMLOptions{Layout: "layout"}
	})
	m.Get("/js/:name.js", rendertext.Renderer(rendertext.Options{
		Directory:  basepath("/view/js"),
		Extensions: []string{".tmpl", ".js"},
		Funcs:      []tt.FuncMap{helpfunc},
		Delims:     rendertext.Delims{"{{", "}}"},
		Charset:    "UTF-8",
		IndentJSON: true,
		IndentXML:  true,
	}), func(params martini.Params, r rendertext.Render) {
		r.Text(200, params["name"], map[string]interface{}{})
	})

	m.Get("/css/:name.css", rendertext.Renderer(rendertext.Options{
		Directory:  basepath("/view/css"),
		Extensions: []string{".tmpl", ".css"},
		Delims:     rendertext.Delims{"{{", "}}"},
		Charset:    "UTF-8",
		IndentJSON: true,
		IndentXML:  true,
	}), func(params martini.Params, r rendertext.Render) {
		r.Text(200, params["name"], map[string]interface{}{})
	})

	m.Get("/conn", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			return
		}
		ag.Join(agent.NewConnWeb(conn))
	})

	m.Get("/conn/:code", func(w http.ResponseWriter, r *http.Request, params martini.Params, session sessions.Session) {
		var user *agent.User
		var conn *HttpConn
		c1, c2, c3, err := rec.Map(params["code"])
		if err != nil {
			return
		}
		b := make(map[string]string, 16)
		r.ParseForm()
		for k, v := range r.Form {
			b[k] = v[0]
		}
		reg := rego.EnJson(b).Bytes()
		reg = append([]byte{0, c1, c2, c3}, reg...)

		if s, ok := session.Get("_id").(uint); ok {
			user = ag.Get(s)
			if user != nil {
				conn, ok = user.Conn.(*HttpConn)
				if !ok {
					return
				}
				conn.Refresh(w, r, reg)
				return
			}
		}
		conn = NewHttpConn(16)
		user = ag.JoinSync(conn)
		session.Set("_id", user.ToUint())
		conn.Refresh(w, r, reg)
	})

	rego.NOTICE("Start from ")
	m.RunOnAddr(fmt.Sprintf(":%d", cfg.Self.ClientPort))
}

func basepath(path string) string {
	return BasePath + path
}
