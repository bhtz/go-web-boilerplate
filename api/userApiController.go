package apiControllers

import (
	"../dal"
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
)

/**
 * HomeController structure.
 */
type UserApiController struct {
	app     *martini.ClassicMartini
	userDal dal.UserDal
}

/**
 * HomeController constructor.
 */
func NewUserApiController(app *martini.ClassicMartini) UserApiController {
	u := dal.NewUserDal()
	a := UserApiController{app, u}
	a.routes()
	return a
}

/**
 * define controller routing.
 */
func (this *UserApiController) routes() {
	this.app.Get("/api/users/self", this.index)
}

/**
 * index action.
 */
func (this *UserApiController) index(r render.Render) {
	user := this.userDal.GetUser()
	r.JSON(200, user)
}
