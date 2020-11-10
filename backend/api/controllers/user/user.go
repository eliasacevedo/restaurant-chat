package user

import (
	"net/http"

	services "fastfoodrestaurant.com/api/services/users"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type UserController struct {
	userService services.IUserService
}

func NewUserController(s services.IUserService) *UserController {
	return &UserController{
		userService: s,
	}
}

const UserBasePath string = "/user"
const userDefaultPath string = "/"
const botDefaultPath string = "/bot"

func (c *UserController) GetRoutes() *chi.Mux {
	userRouter := chi.NewRouter()
	userRouter.Get(userDefaultPath, c.getUser)
	userRouter.Get(botDefaultPath, c.getBot)
	return userRouter
}

func (c *UserController) getUser(rw http.ResponseWriter, r *http.Request) {
	user, _ := c.userService.GetUser()
	render.Render(rw, r, user)
	return
}

func (c *UserController) getBot(rw http.ResponseWriter, r *http.Request) {
	user, _ := c.userService.GetBotInfo()
	render.Render(rw, r, user)
	return
}


