package cyoa

import (
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"
	"strings"
)

var defaultTmpl *template.Template
var defaultTmplHtml = `
<!DOCTYPE html>
<html>
    <head>
        <meta charset="utf-8">
        <title>Choose Your Own Adventure</title>
    </head>
    <body>
		<section class="page">
			<h1>{{.Title}}</h1>
			{{range .Story}}
				<p>{{.}}</p>
			{{end}}
			<u1>
			{{range .Options}}
				<li><a href="/{{.Chapter}}">{{.Text}}</a></li>
			{{end}}    
			</u1>
		</section>
        
		<style>
			body {
				font-family: helvetica, arial;
			}
			h1 {
				text-align:center;
				position:relative;
			}
			.page {
				width: 80%
				max-width: 500px;
				margin: auto;
				margin-top: 40px;
				padding: 80px;
				background: #FFFCF6
				border: 1px solid #eee;
				box-shadow: 0 10px 6px -6px #777;
			}
			ul {
				border-top: 1px dotted #ccc;
				padding: 10px 0 0 0;
				-webkit-padding-start: 0;
			}
			li {
				padding-top: 10px;
			}
			a, 
			a:visited {
				text-decoration: none; 
				color: #6295B5;
			}
			a:active,
			a:hover {
				color: #7792A2;
			}
			p {
				text-indent: 1em;
			}
		</style>
    </body>
</html>
`

func init() {
	defaultTmpl = template.Must(template.New("").Parse(defaultTmplHtml))
}

type HandlerOption func(h *handler)

// functional options https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis
func WithTemplate(t *template.Template) HandlerOption {
	return func(h *handler) {
		h.t = t
	}
}

type handler struct {
	s      Story
	t      *template.Template
	pathFn func(r *http.Request) string
}

func WithPathFn(fn func(r *http.Request) string) HandlerOption {
	return func(h *handler) {
		h.pathFn = fn
	}
}

// if you aren't exporting a type don't return that type explicitly
// it won't export into the docs
// godoc -http :3030
func NewHandler(s Story, opts ...HandlerOption) http.Handler {
	h := handler{
		s,
		defaultTmpl,
		defaultPathFn,
	}
	for _, opt := range opts {
		opt(&h)
	}
	return h
}

func defaultPathFn(r *http.Request) string {
	path := strings.TrimSpace(r.URL.Path)
	if path == "" || path == "/" {
		path = "/intro"
	}
	return path[1:]
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := h.pathFn(r)

	if chapter, ok := h.s[path]; ok {
		err := h.t.Execute(w, chapter)
		if err != nil {
			log.Printf("%v", err)
			http.Error(w, "Something went wrong...", http.StatusInternalServerError)
		}
		return
	}
	http.Error(w, "Chapter not found.", http.StatusNotFound)

}

func JsonStory(r io.Reader) (Story, error) {
	d := json.NewDecoder(r)
	var story Story

	if err := d.Decode(&story); err != nil {
		return nil, err
	}

	return story, nil
}

type Story map[string]Chapter

type Chapter struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}

type Option struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}
