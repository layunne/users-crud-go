package controller

import (
	"net/http"

	"github.com/labstack/echo"
	"gitlab.com/layunne/users-crud-go/models"
	"gitlab.com/layunne/users-crud-go/services"
)

type UsersWebController interface {
	OnGet(echo echo.Context) error
	OnGetAll(echo echo.Context) error
	OnSave(echo echo.Context) error
	OnUpdate(echo echo.Context) error
	OnDelete(echo echo.Context) error
}

func NewUsersWebController(usersService services.UsersService) UsersWebController {
	return &usersWebController{usersService: usersService}
}

type usersWebController struct {
	usersService services.UsersService
}

func (w *usersWebController) OnGet(c echo.Context) error {

	id := c.Param("id")

	user := w.usersService.Get(id)

	if user == nil {
		return c.NoContent(http.StatusNoContent)
	}

	return c.JSON(http.StatusOK, user)
}

func (w *usersWebController) OnGetAll(c echo.Context) error {

	users := w.usersService.GetAll()

	return c.JSON(http.StatusOK, users)
}

func (w *usersWebController) OnSave(c echo.Context) error {

	createUser := &models.CreateUser{}

	if err := c.Bind(createUser); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid payload",
		})
	}

	userResp, err := w.usersService.Save(createUser)

	if err != nil {
		return c.JSON(err.Code, map[string]string{
			"message": err.Info,
		})
	}

	return c.JSON(http.StatusOK, userResp)
}

func (w *usersWebController) OnUpdate(c echo.Context) error {
	user := &models.User{}

	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid payload",
		})
	}

	userResp, err := w.usersService.Update(user)

	if err != nil {
		return c.JSON(err.Code, map[string]string{
			"message": err.Info,
		})
	}

	return c.JSON(http.StatusOK, userResp)
}

func (w *usersWebController) OnDelete(c echo.Context) error {

	id := c.Param("id")

	w.usersService.Delete(id)

	return c.NoContent(http.StatusOK)
}
