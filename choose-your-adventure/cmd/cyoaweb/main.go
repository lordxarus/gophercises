package main

import (
	"flag"
	"fmt"
	cyoa "gophercises/choose-your-adventure/cyoa"
	"log"
	"net/http"
	"os"
)

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

	// tpl := template.Must(template.New("").Parse("Hello"))
	h := cyoa.NewHandler(story)
	fmt.Printf("starting server on port %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
}
