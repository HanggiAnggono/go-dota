package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"hanggi.com/go-dota/handlers"
)

func main() {
	router := gin.New()
	router.Static("/static", "./static")

	router.GET("/", handlers.Home)

  println("Listening on port 3000")
  http.ListenAndServe(":3000", router)
}
