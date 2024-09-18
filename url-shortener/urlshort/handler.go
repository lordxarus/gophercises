package urlshort

import (
	"encoding/json"
	"net/http"

	"gopkg.in/yaml.v2"
)

type pair struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}

// Enum represents the possible options for a certain property.
type DataHandlerType int

const (
	Json DataHandlerType = iota
	Yaml
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if dest, ok := pathsToUrls[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
		} else {
			fallback.ServeHTTP(w, r)
		}
	}
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//   - path: /some-path
//     url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func JsonHandler(in []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var parsed []pair
	err := json.Unmarshal(in, &parsed)
	if err != nil {
		return nil, err
	}
	return MapHandler(pairToMap(parsed), fallback), nil
}

func YamlHandler(in []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var parsed []pair
	err := yaml.Unmarshal(in, &parsed)
	if err != nil {
		return nil, err
	}
	return MapHandler(pairToMap(parsed), fallback), nil
}

func pairToMap(pairs []pair) map[string]string {
	m := make(map[string]string, len(pairs))
	for _, p := range pairs {
		m[p.Path] = p.URL
	}
	return m
}
