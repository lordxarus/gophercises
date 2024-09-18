package main

import (
	"flag"
	"fmt"
	urlshort "gophercises/url-shortener/urlshort"
	"net/http"
	"os"
)

func main() {
	yml := flag.String("yaml", "", "yaml file with urls, an array of entries, each with keys url and path")
	json := flag.String("json", "", "json file with urls, an array of entries, each with keys url and path")
	flag.Parse()

	mux := defaultMux()

	var err error

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	//	// Build the YAMLHandler using the mapHandler as the
	// fallback

	var dataHandler http.HandlerFunc

	switch {
	case *json != "":
		data, err := readFile(*json)
		if err != nil {
			fmt.Printf("error reading file %s: %v\n", *json, err)
		}
		dataHandler, err = urlshort.JsonHandler(data, mapHandler)
		if err != nil {
			fmt.Printf("error creating factory: %v\n", err)
		}
	case *yml != "":
		data, err := readFile(*yml)
		if err != nil {
			fmt.Printf("error reading file %s: %v\n", *yml, err)
		}
		dataHandler, err = urlshort.YamlHandler(data, mapHandler)
		if err != nil {
			fmt.Printf("error creating factory: %v\n", err)
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
	var ret []byte

	file, err := os.Open(path)

	if err != nil {
		return nil, fmt.Errorf("error reading file %s: %v\n", path, err)
	}

	defer file.Close()

	n, err := file.Read(ret)
	if err != nil {
		return nil, fmt.Errorf("error %v reading %s the file. read %d bytes", err, path, n)
	}

	return ret, nil
}
