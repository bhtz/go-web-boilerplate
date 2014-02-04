package controllers

import (
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
)

/**
 * HomeController structure.
 */
type HomeController struct {
	app *martini.ClassicMartini
}

/**
 * HomeController constructor.
 */
func NewHomeController(app *martini.ClassicMartini) HomeController {
	c := HomeController{app}
	c.routes()
	return c
}

/**
 * define controller routing.
 */
func (this *HomeController) routes() {
	this.app.Get("/", this.index)
	this.app.Get("/about", this.about)
	this.app.Get("/home/about", this.about)
	this.app.NotFound(this.httpNotFound)
}

/**
 * index action.
 */
func (this *HomeController) index(r render.Render) {
	r.HTML(200, "home/index", nil)
}

/**
 * about action.
 */
func (this *HomeController) about(r render.Render) {
	r.HTML(200, "home/about", nil)
}

/**
 * httpNotFound action.
 */
func (this *HomeController) httpNotFound() string {
	return "HTTP NOT FOUND !"
}
