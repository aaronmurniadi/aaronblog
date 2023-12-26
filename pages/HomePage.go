package pages

import (
	"github.com/aaronmurniadi/panacea/fragments"
	"github.com/chasefleming/elem-go"
)

func HomePage() string {
	nav := fragments.DefaultNavBar()
	head := fragments.DefaultHead()
	body := elem.Body(nil,
		elem.H1(nil, elem.Text("Hello World")),
	)
	footer := fragments.DefaultFooter()
	return elem.Html(nil, nav, head, body, footer).Render()
}
