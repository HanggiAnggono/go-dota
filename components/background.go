package components

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func Background(children ...Node) Node {
	return Div(Class(" w-full h-full -z-10"),
		Div(Class("fixed top-0 left-0 w-full h-full -z-10 bg-gradient-to-r from-background to-background/30"),
			Div(Class("relative w-full h-full blur-3xl -z-10"),
				Div(Class("animated-background-clip bg-blue-400 w-full h-full absolute top-0 left-0")),
			),
		),
		Group(children),
	)
}
