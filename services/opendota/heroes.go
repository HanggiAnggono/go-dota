package opendota

import (
	"errors"
	"strconv"
	"strings"

	"github.com/samber/lo"
	"hanggi.com/go-dota/api"
)

func GetHeroes() []Hero {
	var heroes []Hero
	api.Fetch().SetResult(&heroes).Get(OpenDotaURL + "/api/heroes")
	constants := GetHeroesConstants()

	for i := range heroes {
		hero := &heroes[i]
		id := strconv.Itoa(hero.ID)
		hero.Img = constants[id].Img
	}

	return heroes
}

func GetHeroDetail(name string) (HeroDetail, error) {
	var heroDetail HeroDetail
	heroes := GetHeroes()
	hero, found := lo.Find(heroes, func(hero Hero) bool {
		return hero.Name == name
	})

	if !found {
		return heroDetail, errors.New("Hero not found")
	}

	heroDetail.Hero = hero
	heroDetail.Constants = GetHeroesConstants()[name]
	heroDetail.Lore = GetHeroLore(name)
	heroDetail.HeroAbilities = GetHeroAbilities(name)
	return heroDetail, nil
}

func GetHeroLore(name string) string {
	shortname := strings.Replace(name, "npc_dota_hero_", "", 1)
	if len(HeroesLores) > 0 {
		return HeroesLores[shortname]
	}
	api.Fetch().SetResult(&HeroesLores).Get(OpenDotaURL + "/api/constants/hero_lore")
	return HeroesLores[shortname]
}

func GetHeroesConstants() HeroesConstantsMap {
	if len(HeroesConstants) > 0 {
		return HeroesConstants
	}

	api.Fetch().SetResult(&HeroesConstants).Get(OpenDotaURL + "/api/constants/heroes")
	return HeroesConstants
}

func GetHeroAbilities(name string) HeroAbilities {
	if len(HeroAbilitiesData) > 0 {
		return HeroAbilitiesData[name]
	}

	api.Fetch().SetResult(&HeroAbilitiesData).Get(OpenDotaURL + "/api/constants/hero_abilities")
	return HeroAbilitiesData[name]
}

func GetAbilities() map[string]Ability {
	if len(AbilitiesData) > 0 {
		return AbilitiesData
	}

	api.Fetch().SetResult(&AbilitiesData).Get(OpenDotaURL + "/api/constants/abilities")
	return AbilitiesData
}

var HeroesConstants HeroesConstantsMap = map[string]HeroConstant{}
var HeroesLores = map[string]string{}
var HeroAbilitiesData = map[string]HeroAbilities{}
var AbilitiesData = map[string]Ability{}

type Hero struct {
	ID            int      `json:"id"`
	Name          string   `json:"name"`
	LocalizedName string   `json:"localized_name"`
	PrimaryAttr   string   `json:"primary_attr"`
	AttackType    string   `json:"attack_type"`
	Roles         []string `json:"roles"`
	Legs          int      `json:"legs"`
	Img           string   `json:"icon"`
}

type HeroDetail struct {
	Hero
	Constants     HeroConstant  `json:"constants"`
	Lore          string        `json:"lore"`
	HeroAbilities HeroAbilities `json:"abilities"`
}

type HeroesConstantsMap = map[string]HeroConstant

type HeroConstant struct {
	ID              int      `json:"id"`
	Name            string   `json:"name"`
	PrimaryAttr     string   `json:"primary_attr"`
	AttackType      string   `json:"attack_type"`
	Roles           []string `json:"roles"`
	Img             string   `json:"img"`
	Icon            string   `json:"icon"`
	BaseHealth      int      `json:"base_health"`
	BaseHealthRegen int      `json:"base_health_regen"`
	BaseMana        int      `json:"base_mana"`
	BaseManaRegen   int      `json:"base_mana_regen"`
	BaseArmor       int      `json:"base_armor"`
	BaseMR          int      `json:"base_mr"`
	BaseAttackMin   int      `json:"base_attack_min"`
	BaseAttackMax   int      `json:"base_attack_max"`
	BaseSTR         int      `json:"base_str"`
	BaseAGI         int      `json:"base_agi"`
	BaseINT         int      `json:"base_int"`
	STRGain         float64  `json:"str_gain"`
	AGIGain         float64  `json:"agi_gain"`
	INTGain         float64  `json:"int_gain"`
	AttackRange     int      `json:"attack_range"`
	ProjectileSpeed int      `json:"projectile_speed"`
	AttackRate      float64  `json:"attack_rate"`
	BaseAttackTime  int      `json:"base_attack_time"`
	AttackPoint     float64  `json:"attack_point"`
	MoveSpeed       int      `json:"move_speed"`
	TurnRate        *float64 `json:"turn_rate"`
	CMEnabled       bool     `json:"cm_enabled"`
	Legs            int      `json:"legs"`
	DayVision       int      `json:"day_vision"`
	NightVision     int      `json:"night_vision"`
	LocalizedName   string   `json:"localized_name"`
}

type HeroAbilities struct {
	Abilities []string `json:"abilities"`
	Talents   []struct {
		Name  string `json:"name"`
		Level int    `json:"level"`
	} `json:"talents"`
	Facets []struct {
		ID          int    `json:"id"`
		Name        string `json:"name"`
		Icon        string `json:"icon"`
		Color       string `json:"color"`
		GradientID  int    `json:"gradient_id"`
		Title       string `json:"title"`
		Description string `json:"description"`
	} `json:"facets"`
}

type Ability struct {
	Dname     string   `json:"dname"`
	Behavior  string   `json:"behavior"`
	DmgType   string   `json:"dmg_type"`
	BkbPierce string   `json:"bkbpierce"`
	Desc      string   `json:"desc"`
	Attrib    []Attrib `json:"attrib"`
	Lore      string   `json:"lore"`
	Img       string   `json:"img"`
	IsInnate  bool     `json:"is_innate"`
	CD        any      `json:"cd"` // cooldown
	MC        any      `json:"mc"` // manacost
}

func (a *Ability) GetCooldownString() string {
	var cd string
	switch a.CD.(type) {
	case string:
		cd = a.CD.(string)
	case []string:
		cd = strings.Join(a.CD.([]string), "/")
	}
	return cd
}

func (a *Ability) GetManacostString() string {
	var mc string
	switch a.MC.(type) {
	case string:
		mc = a.MC.(string)
	case []string:
		mc = strings.Join(a.MC.([]string), "/")
	}
	return mc
}

type Attrib struct {
	Key       string      `json:"key"`
	Header    string      `json:"header"`
	Value     interface{} `json:"value"`
	Generated bool        `json:"generated,omitempty"`
}
