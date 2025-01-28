package components

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func Background(children ...Node) Node {
	return Div(Class(" w-full h-full -z-10"),
		Div(Class("fixed top-0 left-0 w-full h-full -z-10 bg-gradient-to-r from-black to-gray-700/80"),
			Div(Class("relative w-full h-full blur-3xl -z-10"),
				Div(Class("animated-background-clip bg-blue-400 w-full h-full absolute top-0 left-0")),
			),
			Div(Class("absolute right-[10rem] top-[10rem] w-[15rem] h-[120vh] rotate-[45deg] bg-gradient-to-b from-black/40 to-transparent")),
		),
		Group(children),
	)
}
