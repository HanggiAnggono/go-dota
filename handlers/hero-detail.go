package handlers

import (
	"github.com/gin-gonic/gin"
	"hanggi.com/go-dota/pages"
	"hanggi.com/go-dota/services/opendota"
)

func HeroDetail(c *gin.Context) {
	name := c.Param("name")
	heroDetail, err := opendota.GetHeroDetail(name)

	if err != nil {
		c.String(404, "Hero not found")
	}

	pages.HeroDetail(&heroDetail).Render(c.Writer)
}
