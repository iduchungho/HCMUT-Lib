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
func (app *App) Init() {
	app.r = gin.Default()
}

func (app *App) Run() {
	if app.r != nil {
		// route service
		route.UserRoutes(app.r)
		app.r.Run("localhost:8080")
	} else {
		panic("gin engine is not available")
	}
}
