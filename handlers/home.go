package handlers

import (
	"github.com/gin-gonic/gin"
	"hanggi.com/go-dota/pages"
	"hanggi.com/go-dota/services/opendota"
)

func Home(c *gin.Context) {
	heroes := opendota.GetHeroes()
	pages.Home(&heroes).Render(c.Writer)
}
