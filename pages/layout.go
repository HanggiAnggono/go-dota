package pages

import (
	"github.com/samber/lo"
	"hanggi.com/go-dota/components"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func Layout(title string, children ...Node) Node {
	pageTitle := lo.If(len(title) > 0, title).Else("Go Dota")
	return HTML(
		Head(
			Link(Rel("icon"), Type("image/png"), Href("/static/favicon.png")),
			Title(pageTitle),
			TitleEl(Text(pageTitle)),
			Tailwind(),
		),
		Body(
			Div(
				Class("min-h-screen"),
				components.Background(Group(children)),
			),
		),
	)
}
