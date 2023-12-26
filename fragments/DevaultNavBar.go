package fragments

import (
	"github.com/chasefleming/elem-go"
	"github.com/chasefleming/elem-go/attrs"
)

func DefaultNavBar() *elem.Element {
	nameLinkMap := map[string]string{
		"Home":  "/",
		"About": "/about",
		"Posts": "/posts",
	}
	table := elem.Table(attrs.Props{attrs.Class: "navbar"})
	tr := elem.Tr(nil)
	for name, path := range nameLinkMap {
		tdElement := elem.Td(
			attrs.Props{attrs.Style: "text-align:center; border: 0px"},
			elem.A(attrs.Props{attrs.Href: path},
				elem.Text(name)))
		tr.Children = append(tr.Children, tdElement)
	}
	table.Children = append(table.Children, tr)
	return table
}
