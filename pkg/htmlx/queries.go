package htmlx

import (
	"golang.org/x/net/html"
)

func GetElementByClassname(n *html.Node, classes string) *html.Node {
	return traverseFirst(n, tagHasClasses(classes))
}

func GetElementsByClassname(n *html.Node, classes string) (nodes []*html.Node) {
	traverse(n, tagHasClasses(classes), &nodes, -1)
	return
}

func GetElementById(n *html.Node, id string) *html.Node {
	return traverseFirst(n, tagHasAttr("id", id))
}

func GetElementsByTagName(n *html.Node, tag string) (nodes []*html.Node) {
	traverse(n, tagIs(tag), &nodes, -1)
	return
}

func GetElementByTagName(n *html.Node, tag string) *html.Node {
	return traverseFirst(n, tagIs(tag))
}

func GetElementByAttribute(n *html.Node, attr string, value string) *html.Node {
	return traverseFirst(n, tagHasAttr(attr, value))
}
