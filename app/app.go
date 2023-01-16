package app

import (
	"example/hcmut-lib/config"
	"example/hcmut-lib/pkg/route"

	"github.com/gin-gonic/gin"
)

type App struct {
	r *gin.Engine
}

// app constructor
func init() {
	config.DatabaseConection()
}

// service method
func (app App) Run() {

	app.r = gin.Default()

	// route service
	route.UserRoutes(app.r)

	app.r.Run("localhost:8080")
}
