package pages

import (
	"strings"

	"github.com/samber/lo"
	"hanggi.com/go-dota/services/opendota"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func Home(heroes *[]opendota.Hero) Node {
	attributes := lo.Uniq(lo.Map(*heroes, func(hero opendota.Hero, i int) string {
		return hero.PrimaryAttr
	}))
	heroesByAttributes := lo.GroupBy(*heroes, func(hero opendota.Hero) string {
		return hero.PrimaryAttr
	})

	return Layout(
		Div(
			Class("flex flex-col items-center justify-center p-10"),
			H1(Class("text-5xl font-extrabold bg-gradient-to-r from-blue-500 to-pink-500 bg-clip-text text-transparent"), Text("Welcome to Go-Dota")),
			P(Text("Build your custom hero with this tool")),

			Div(
				Class("grid grid-cols-2 gap-10"),
				Map(attributes, func(attr string) Node {
					return Div(
						Class("mb-10"),
						H1(Class("text-2xl mb-4"), Text(strings.ToUpper(attr))),
						Div(
							Class("grid grid-cols-7 gap-1"),
							Map(heroesByAttributes[attr], func(hero opendota.Hero) Node {

								return Div(
									Img(Class("w-20 h-14"),Src(opendota.DotaImageHost+hero.Img)),
								)
							}),
						),
					)
				}),
			),
		),
	)
}
