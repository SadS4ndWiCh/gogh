package htmlx

import (
	"bytes"
	"io"
	"log"
	"strings"

	"golang.org/x/net/html"
)

func Load(content []byte) (*html.Node, error) {
	return html.Parse(strings.NewReader(string(content)))
}

func traverse(n *html.Node, cond func(*html.Node) bool, dst *[]*html.Node, depth int) {
	if cond(n) {
		*dst = append(*dst, n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if len(*dst) == depth {
			return
		}
		traverse(c, cond, dst, depth)
	}
}

func traverseFirst(n *html.Node, cond func(*html.Node) bool) *html.Node {
	if cond(n) {
		return n
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if _n := traverseFirst(c, cond); _n != nil {
			return _n
		}
	}

	return nil
}

func RenderNode(n *html.Node) string {
	var buf bytes.Buffer
	w := io.Writer(&buf)

	if err := html.Render(w, n); err != nil {
		log.Println("Failed to render node: ", err)
		return ""
	}

	return buf.String()
}
