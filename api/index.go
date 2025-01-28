package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"hanggi.com/go-dota/handlers"
)

var (
	app *gin.Engine
)

// @title Golang Vercel Deployment
// @description API Documentation for Golang deployment in vercel serverless environment
// @version 1.0

// @schemes https http
// @host golang-vercel.vercel.app
func init() {
	app = gin.New()
	handlers.Register(app)
}

// Entrypoint
func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}
