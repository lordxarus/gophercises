package link

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

// Link represents a link (<a href="...">) in an HTML document in
// an HTML document
type Link struct {
	Href string
	Text string
}

// Parse will take in an HTML document and will return a slice of
// links parsed from it.
func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	nodes := searchLinks(doc)
	var links []Link
	for _, node := range nodes {
		links = append(links, buildLink(node))
	}
	return links, nil
}

// build a Link struct from an html.Node
func buildLink(n *html.Node) Link {
	var ret Link
	for _, attr := range n.Attr {
		if attr.Key == "href" {
			ret.Href = attr.Val
			break
		}
	}
	ret.Text = strings.Join(strings.Fields(searchText(n)), " ")
	return ret
}

func searchText(n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	}
	if n.Type != html.ElementNode {
		return ""
	}
	var ret string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret += searchText(c)
	}
	return ret
}

// dfs
func searchLinks(n *html.Node) []*html.Node {
	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}
	var ret []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret = append(ret, searchLinks(c)...)
	}
	return ret
}
