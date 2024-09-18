package main

import (
	"flag"
	"fmt"
	urlshort "gophercises/url-shortener/urlshort"
	"net/http"
	"os"
	"strings"
)

func main() {
	var path string
	flag.StringVar(&path, "file", "", "input urls. can be in json or yaml, an array of entries, each with keys url and path")
	flag.Parse()

	mux := defaultMux()

	var err error

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	var dataHandler http.HandlerFunc

	if path != "" {
		data, err := readFile(path)
		if err != nil {
			fmt.Printf("error reading file %s: %v\n", path, err)
		}

		spl := strings.Split(path, ".")
		ext := spl[len(spl)-1]

		switch ext {
		case "json":
			fmt.Println("Json")
			dataHandler, err = urlshort.JsonHandler(data, mapHandler)
		case "yml", "yaml":
			fmt.Println("Yaml")
			dataHandler, err = urlshort.YamlHandler(data, mapHandler)
		}

		if err != nil {
			panic(fmt.Sprintf("error creating handler: %v\n", err))
		}

	}

	fmt.Println("Starting the server on :8080")
	err = http.ListenAndServe(":8080", dataHandler)
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

func readFile(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("%v while reading file %s", path, err)
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		return nil, fmt.Errorf("couldn't get file metadata %v", err)
	}

	// if you don't set the initial size you get
	// unexpected end of JSON input
	// https://stackoverflow.com/a/65414065
	ret := make([]byte, info.Size())

	n, err := file.Read(ret)
	if err != nil {
		return nil, fmt.Errorf("%v while reading %s. read %d bytes", err, path, n)
	}

	return ret, nil
}
