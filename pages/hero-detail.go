package pages

import (
	"fmt"
	"strings"

	"github.com/samber/lo"
	"hanggi.com/go-dota/components"
	"hanggi.com/go-dota/services/opendota"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func HeroDetail(heroDetail *opendota.HeroDetail) Node {
	shortname := strings.Replace(heroDetail.Name, "npc_dota_hero_", "", 1)

	return Layout(
		fmt.Sprintf("Go Dota | %s", heroDetail.LocalizedName),
		HeroDetailHeader(heroDetail),
		HeroDetailAbilities(shortname, heroDetail.HeroAbilities),
	)
}

func HeroDetailHeader(heroDetail *opendota.HeroDetail) Node {
	shortname := strings.Replace(heroDetail.Name, "npc_dota_hero_", "", 1)

	return Div(
		Class("flex"),
		Div(
			Class("p-10 flex flex-col w-2/5"),
			H1(
				Class("text-5xl font-extrabold bg-gradient-to-r from-blue-500 to-pink-500 bg-clip-text text-transparent"),
				Text(heroDetail.LocalizedName),
			),
			Details(
				Summary(
					Class("text-lg mt-4"),
					Text(heroDetail.Lore[:100]+"...Click to read more"),
				),
				Span(
					Text(heroDetail.Lore),
				),
			),
		),
		Div(Class("flex w-3/5"),
			Div(
				Class("relative"),
				components.VideoAsset(fmt.Sprintf("heroes/renders/%s", shortname)),
				Div(
					Class("absolute bottom-0 right-[25%] flex gap-2"),
					Map(heroDetail.HeroAbilities.Abilities, func(ability string) Node {
						abilityDetail := opendota.GetAbilities()[ability]
						return Div(
							Div(
								If(!abilityDetail.IsInnate, Div(
									Img(Class("size-20 rounded-md"), Src(opendota.DotaImageHost+abilityDetail.Img)),
								)),
							),
						)
					}),
				),
			),
		),
	)
}

func HeroDetailAbilities(heroShortname string, heroAbilities opendota.HeroAbilities) Node {
	abilites := opendota.GetAbilities()
	heroSkills := lo.Filter(heroAbilities.Abilities, func(ability string, i int) bool {
		return !abilites[ability].IsInnate
	})

	return Div(
		Class("mt-[40vw] p-10 relative"),
		Div(
			Class("flex gap-2"),
			Map(heroSkills, func(ability string) Node {
				index := lo.IndexOf(heroAbilities.Abilities, ability)
				abilityDetail := abilites[ability]

				return Div(
					Class("flex flex-col-reverse"),
					Label(
						Class("peer"),
						Input(
							Class("invisible peer/a"),
							Type("radio"),
							Name("ability"),
							ID(ability),
							If(index == 0, Checked()),
						),
						Img(
							Class("size-20 rounded-md peer-checked/a:grayscale-0 grayscale"),
							Src(opendota.DotaImageHost+abilityDetail.Img),
						),
					),
					Div(
						Class("peer-has-[:checked]:opacity-100 opacity-0 transition-opacity absolute left-[2rem] bottom-[5rem] w-[90vw] -z-10"),
						Div(
							Class("flex gap-4"),
							Div(
								Class("w-2/3"),
								components.BorderGradientBox(
									components.VideoAsset(
										fmt.Sprintf("abilities/%s/%s", heroShortname, ability),
										Class("rounded-md"),
									),
								),
							),
							Div(
								Class("bg-background p-2 w-1/3 flex gap-2"),
								Img(
									Class("size-20"),
									Src(opendota.DotaImageHost+abilityDetail.Img),
									Alt(abilityDetail.Dname),
								),
								Div(
									H2(Class("text-2xl font-bold"), Text(abilityDetail.Dname)),
									P(Text(lo.If(len(abilityDetail.Desc) > 0, abilityDetail.Desc).Else(abilityDetail.Lore))),
								),
							),
						),
					),
				)
			}),
		),
	)
}
