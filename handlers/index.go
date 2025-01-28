package handlers

import "github.com/gin-gonic/gin"

func Register(app *gin.Engine) {
	app.Static("/static", "./static")

	app.GET("/hero/:name", HeroDetail)
	app.GET("/", Home)
}