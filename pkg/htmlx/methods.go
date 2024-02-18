package htmlx

import (
	"strings"

	"golang.org/x/net/html"
)

func GetAttribute(n *html.Node, key string) (string, bool) {
	for _, attr := range n.Attr {
		if attr.Key == key {
			return attr.Val, true
		}
	}

	return "", false
}

func GetTextContent(n *html.Node) (string, bool) {
	if n == nil {
		return "", false
	}

	if n.Type == html.TextNode {
		return strings.TrimSpace(n.Data), true
	}

	return GetTextContent(n.FirstChild)
}
