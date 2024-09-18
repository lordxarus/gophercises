package main

import (
	"flag"
	"fmt"
	cyoa "gophercises/choose-your-adventure/cyoa"
	"html/template"
	"log"
	"net/http"
	"os"
)

var storyTmpl = `
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
				<li><a href="/story/{{.Chapter}}">{{.Text}}</a></li>
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

func main() {
	port := flag.Int("port", 3000, "port for cyoaweb server")
	path := flag.String("file", "gopher.json", "the JSON file with the CYOA story")
	flag.Parse()
	fmt.Printf("using %s\n", *path)

	f, err := os.Open(*path)
	if err != nil {
		panic(err)
	}

	story, err := cyoa.JsonStory(f)

	if err != nil {
		panic(err)
	}

	tpl := template.Must(template.New("").Parse(storyTmpl))
	h := cyoa.NewHandler(story,
		cyoa.WithTemplate(tpl),
		cyoa.WithPathFn(pathFn),
	)
	mux := http.NewServeMux()
	mux.Handle("/story/", h)

	fmt.Printf("starting server on port %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), mux))
}

func pathFn(r *http.Request) string {
	path := r.URL.Path
	if path == "/story" || path == "/story/" {
		path = "/story/intro"
	}
	return path[len("/story/"):]
}
