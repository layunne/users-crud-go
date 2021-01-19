package application

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gitlab.com/layunne/users-crud-go/config"
	"gitlab.com/layunne/users-crud-go/controller"
)

func NewUsersWebServer(env config.Env, controller controller.UsersWebController) Server {
	return &webServer{controller:controller, port:env.WebServerPort()}
}

type webServer struct {
	controller controller.UsersWebController
	port       string
}

func (w *webServer) Start() error {

	e := echo.New()

	e.Use(middleware.Recover())

	e.GET("/users/:id", 	 func(context echo.Context) error { return w.controller.OnGet(context) })
	e.GET("/users",     	 func(context echo.Context) error { return w.controller.OnGetAll(context) })
	e.POST("/users",    	 func(context echo.Context) error { return w.controller.OnSave(context) })
	e.PUT("/users",        func(context echo.Context) error { return w.controller.OnUpdate(context) })
	e.DELETE("/users/:id", func(context echo.Context) error { return w.controller.OnDelete(context) })

	return e.Start(fmt.Sprintf(":%s", w.port))
}
