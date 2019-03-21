package vugu

import (
	"strings"

	"golang.org/x/net/html"
)

// stuff that is common to both parsers can get moved into here

func staticVGAttr(inAttr []html.Attribute) (ret []VGAttribute) {

	for _, a := range inAttr {
		switch {
		case a.Key == "vg-if":
		case a.Key == "vg-for":
		case a.Key == "vg-html":
		case strings.HasPrefix(a.Key, ":"):
		case strings.HasPrefix(a.Key, "@"):
		default:
			ret = append(ret, attrFromHtml(a))
		}
	}

	return ret
}

func vgIfExpr(n *html.Node) string {
	for _, a := range n.Attr {
		if a.Key == "vg-if" {
			return a.Val
		}
	}
	return ""
}

func vgForExpr(n *html.Node) string {
	for _, a := range n.Attr {
		if a.Key == "vg-for" {

			v := strings.TrimSpace(a.Val)

			if !strings.Contains(v, " ") { // make it so `w` is a shorthand for `key, value := range w`
				v = "key, value := range " + v
			}

			return v
		}
	}
	return ""
}

func vgHTMLExpr(n *html.Node) string {
	for _, a := range n.Attr {
		if a.Key == "vg-html" {
			return a.Val
		}
	}
	return ""
}