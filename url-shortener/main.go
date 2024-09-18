package main

import (
	"flag"
	"fmt"
	urlshort "gophercises/url-shortener/urlshort"
	"net/http"
	"os"
)

func main() {
	yml := flag.String("yaml", "urls.yml", "yaml file with urls, an array of entries, each with keys url and path")
	flag.Parse()

	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	//	// Build the YAMLHandler using the mapHandler as the
	// fallback

	file, err := os.Open(*yml)
	if err != nil {
		fmt.Printf("Error opening the file %s: %v\n", *yml, err)
	}
	defer file.Close()

	var ymlData []byte
	n, err := file.Read(ymlData)
	if err != nil {
		fmt.Printf("Error: %v reading %s the file. Read %d bytes.", err, *yml, n)
	}
	yamlHandler, err := urlshort.YAMLHandler(ymlData, mapHandler)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on :8080")
	err = http.ListenAndServe(":8080", yamlHandler)
	if err != nil {
		panic(err)
	}
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
