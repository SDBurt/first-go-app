package handler

import (
	"first-go/model"
	"first-go/view/user"

	"github.com/labstack/echo/v4"
)

type UserHandler struct{}

func (h UserHandler) HandleUserShow(c echo.Context) error {

	u := model.User{
		Email: "test@test.ca",
	}

	return render(c, user.Show(u))
}
