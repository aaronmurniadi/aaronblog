package fragments

import (
	"github.com/chasefleming/elem-go"
	"github.com/chasefleming/elem-go/attrs"
)

func DefaultHead() *elem.Element {
	return elem.Head(nil,
		elem.Script(attrs.Props{
			attrs.Src:         "https://unpkg.com/htmx.org@1.9.6",
			attrs.Integrity:   "sha384-FhXw7b6AlE/jyjlZH5iHa/tTe9EpJ1Y55RjcgPbjeWMskSxZt1v9qkxLJWNJaGni",
			attrs.Crossorigin: "anonymous"}),
		elem.Link(attrs.Props{
			attrs.Rel:  "stylesheet",
			attrs.Href: "https://cdn.simplecss.org/simple.min.css"}))
}
