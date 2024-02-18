package htmlx

import (
	"golang.org/x/net/html"
)

func tagHasClasses(classes string) func(*html.Node) bool {
	return func(n *html.Node) bool {
		if n.Type != html.ElementNode {
			return false
		}

		classnames, exists := GetAttribute(n, "class")
		if !exists {
			return false
		}

		return stringIncludesWords(classnames, classes)
	}
}

func tagHasAttr(attr string, value string) func(*html.Node) bool {
	return func(n *html.Node) bool {
		attrValue, exists := GetAttribute(n, attr)

		return exists && attrValue == value
	}
}

func tagIs(tag string) func(*html.Node) bool {
	return func(n *html.Node) bool {
		return n.Data == tag
	}
}
