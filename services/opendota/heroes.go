package opendota

import (
	"strconv"

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

func GetHeroesConstants() HeroesConstantsMap {
	if len(HeroesConstants) > 0 {
		return HeroesConstants
	}

	api.Fetch().SetResult(&HeroesConstants).Get(OpenDotaURL + "/api/constants/heroes")
	return HeroesConstants
}

var HeroesConstants HeroesConstantsMap = map[string]HeroConstant{}

type Hero struct {
	ID            int      `json:"id"`
	Name          string   `json:"name"`
	LocalizedName string   `json:"localized_name"`
	PrimaryAttr   string   `json:"primary_attr"`
	AttackType    string   `json:"attack_type"`
	Roles         []string `json:"roles"`
	Legs          int      `json:"legs"`
	Img          string   `json:"icon"`
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
