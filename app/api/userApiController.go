package apiControllers

import (
	"../dal"
	"../models"
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
	"net/http"
	"strconv"
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
	this.app.Get("/api/users/self", this.connected)
	this.app.Get("/api/users/:id", this.get)
	this.app.Get("/api/users", this.getAll)
}

/**
 * getAll action.
 */
func (this *UserApiController) getAll(r render.Render) {
	r.JSON(200, this.userDal.GetAll())
}

/**
 * get action.
 */
func (this *UserApiController) get(w http.ResponseWriter, params martini.Params, r render.Render) {
	var user models.User
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user = this.userDal.Get(id)
	r.JSON(200, user)
}

/**
 * get connected user action.
 */
func (this *UserApiController) connected(r render.Render) {
	user := this.userDal.GetUser()
	r.JSON(200, user)
}
