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
	if n.FirstChild == nil {
		return "", false
	}

	if n.FirstChild.Type == html.TextNode {
		return strings.TrimSpace(n.FirstChild.Data), true
	}

	return GetTextContent(n.FirstChild)
}
