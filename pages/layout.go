package pages

import (
	"hanggi.com/go-dota/components"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func Layout(children ...Node) Node {
	return HTML(
		Head(
			Title("Go Dota"),
			TitleEl(Text("Go Dota")),
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
