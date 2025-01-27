package pages

import (
	"os"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

// TODO: handle dark and light mode
var tailwindConfig, err1 = os.ReadFile("static/tailwind.config.js")
var tailwindConfigString = string(tailwindConfig)

var globalsCSS, err = os.ReadFile("static/globals.css")
var globalsCSSString = string(globalsCSS)


func Tailwind() Node {
	if err1 != nil {
		println(err1.Error())
	}

	if err != nil {
		println(err.Error())
	}

	return Group(
		[]Node{
			Script(Src("https://cdn.tailwindcss.com")),
			Script(Raw(tailwindConfigString)),
			StyleEl(Type("text/tailwindcss"), Raw(globalsCSSString)),
		},
	)
}
