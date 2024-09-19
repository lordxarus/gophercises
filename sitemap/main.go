package main

import (
	"flag"
	"fmt"
	"gophercises/html-link-parser/test-html/link"
	"io"
	"net/http"
	"net/url"
	"strings"
)

/*
1. GET the webpage
2. parse all the links on the page
3. build proper urls with our links
4. filter out any links with a different domain
5. find all pages (BFS)
6. print out XML
*/
func main() {
	urlFlag := flag.String("url", "https://gophercises.com", "the url that you want to build a sitemap for")
	flag.Parse()

	fmt.Println(*urlFlag)
	resp, err := http.Get(*urlFlag)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	/*
	 /some-path
	 https://gophercises.com/some-path
	 #fragment
	 mailto:seraphim@jaz.codes
	*/

	// We are using the url from our request because we might have been redirected
	reqUrl := resp.Request.URL
	baseUrl := &url.URL{
		Scheme: reqUrl.Scheme,
		Host:   reqUrl.Host,
	}
	base := baseUrl.String()

	for _, href := range hrefs {
		fmt.Println(href)
	}

}

func hrefs(r io.Reader, base string) []string {
	links, _ := link.Parse(r)
	var hrefs []string
	for _, l := range links {
		switch {
		case strings.HasPrefix(l.Href, "/"):
			hrefs = append(hrefs, base+l.Href)
		case strings.HasPrefix(l.Href, "http"):
			hrefs = append(hrefs, l.Href)
		}
	}
}
