package users

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
)

type createUserFn func(context.Context, *Users) error

func (fn createUserFn) createUserFn(ctx context.Context, user *Users) error {
	return fn(ctx, user)
}

func CreateUserHandler(scu createUserFn) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := new(Users)
		if err := c.Bind(user); err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		if err := scu.createUserFn(c.Request().Context(), user); err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, "user created")

	}
}
