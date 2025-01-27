package components

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func BorderGradientBox(children ...Node) Node {
	return Div(
		Class("[background-image:conic-gradient(cyan,white,cyan)] relative p-[1px] rounded-md"),
		Div(Class("absolute top-[50%] left-[50%] translate-x-[-50%] translate-y-[-50%] w-full h-full blur-3xl rounded-md [background-image:conic-gradient(cyan,white,cyan)] -z-10")),
		Group(children),
	)
}
