package main

import (
	"example/hcmut-lib/app"
)

func main() {
	var app = new(app.App) // create a new app
	app.Init()             // initialize app
	app.Run()              // run the app
}
