package pages

import (
	"encoding/json"
	"fmt"
	"os"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

// TODO: handle dark and light mode
var colors = map[string]string{
	"background": "#3e3a44",
	"foreground": "#d3d3d3",
	"primary":    "#1E40AF",
	"secondary":  "#1F2937",
}

var jsonColors, _ = json.Marshal(colors)

var globalsCSS, err = os.ReadFile("static/globals.css")
var globalsCSSString = string(globalsCSS)


func Tailwind() Node {
	if err != nil {
		println(err.Error())
	}

	return Group(
		[]Node{
			Script(Src("https://cdn.tailwindcss.com")),
			Script(Raw(
				`tailwind.config = {
				theme: {
					extend: {
						colors: ` + fmt.Sprintf("%v", string(jsonColors)) + `,
					}
				}
			}`,
			)),
			StyleEl(Type("text/tailwindcss"), Raw(globalsCSSString)),
		},
	)
}
