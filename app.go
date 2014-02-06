package main

import (
	"./api"
	"./controllers"
	"./db"
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
)

/**
 * Connect to database.
 */
func dbConnection() {
	db.Init()
	db.CreateTables()
}

/**
 * Load all Controllers.
 */
func loadControllers(app *martini.ClassicMartini) {
	controllers.NewHomeController(app)
}

/**
 * Load all ApiControllers.
 */
func loadApiControllers(app *martini.ClassicMartini) {
	apiControllers.NewUserApiController(app)
}

/**
 * Martini application configuration.
 */
func configuration(app *martini.ClassicMartini) {
	app.Use(martini.Static("public"))
	app.Use(render.Renderer(render.Options{
		Directory:  "views",
		Layout:     "layout",
		Charset:    "UTF-8",
		IndentJSON: true,
	}))
}

/**
 * Run martini application.
 */
func main() {
	app := martini.Classic()
	dbConnection()
	configuration(app)
	loadControllers(app)
	loadApiControllers(app)
	app.Run()
}
