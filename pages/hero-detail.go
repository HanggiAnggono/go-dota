package pages

import (
	"fmt"
	"strings"

	"github.com/samber/lo"
	"hanggi.com/go-dota/components"
	"hanggi.com/go-dota/services/opendota"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/components"
	. "maragu.dev/gomponents/html"
)

func HeroDetail(heroDetail *opendota.HeroDetail) Node {
	shortname := strings.Replace(heroDetail.Name, "npc_dota_hero_", "", 1)

	return Layout(
		fmt.Sprintf("Go Dota | %s", heroDetail.LocalizedName),
		HeroDetailHeader(heroDetail),
		HeroDetailAttributes(heroDetail),
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
				Class("text-8xl font-extrabold bg-gradient-to-r from-blue-500 to-pink-500 bg-clip-text text-transparent"),
				Text(heroDetail.LocalizedName),
			),
			Details(
				Summary(
					Class("text-lg mt-4"),
					Text(heroDetail.Lore[:200]+"...Click to read more"),
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
									Img(Class("size-20 rounded-md shadow-md relative z-20"), Src(opendota.DotaImageHost+abilityDetail.Img)),
								)),
							),
						)
					}),
				),
			),
		),
	)
}

func HeroDetailAttributes(heroDetail *opendota.HeroDetail) Node {
	return Div(
		Class("relative z-10 h-40 mt-4 mb-10 [box-shadow:black_0rem_-5rem_8rem_8rem] bg-gradient-to-l from-gray-600 to-black p-10 flex items-center justify-center"),
		Div(
			Img(Src(opendota.DotaImageHost+heroDetail.Img)),
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
						Class("peer shadow-md hover:scale-110 cursor-pointer transition-transform"),
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
							AbilityDetail(abilityDetail),
						),
					),
				)
			}),
		),
	)
}

func AbilityDetail(abilityDetail opendota.Ability) Node {
	return Div(
		Class("bg-gray-500 flex flex-col w-1/3 max-h-[720px] overflow-visible"),
		Div(
			Class("flex p-4 gap-2"),
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
		Div(
			Class("bg-gray-800 p-4 flex-grow"),
			Map(abilityDetail.Attrib, func(attrib opendota.Attrib) Node {
				var values string
				switch attrib.Value.(type) {
				case string:
					values = attrib.Value.(string)
				case []string:
					values = strings.Join(attrib.Value.([]string), "/")
				}

				return Div(
					Class("flex gap-2"),
					Span(Class("text-gray-500"), Text(strings.Title(strings.ToLower(attrib.Header)))),
					Span(Text(values)),
				)
			}),
			Div(
				Class("mt-auto flex justify-between"),
				// cooldown
				Div(
					Class("flex items-center gap-2"),
					Div(Classes{
						"size-4 rounded-sm [background-image:conic-gradient(grey_0%,grey_40%,black_50%,black_100%)]": true,
						"invisible": len(abilityDetail.GetCooldownString()) == 0,
					}),
					Span(Text(abilityDetail.GetCooldownString())),
				),
				Div(
					Class("flex items-center gap-2"),
					Div(Class("size-4 bg-blue-400 rounded-sm")),
					Span(Text(abilityDetail.GetManacostString())),
				),
			),
		),
	)
}
