package users

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
)

type getAllusersFn func(context.Context) ([]*Users, error)

func (gn getAllusersFn) GetAllUser(ctx context.Context) ([]*Users, error) {
	return gn(ctx)
}

func GetAllUserHandler(svc getAllusersFn) echo.HandlerFunc {
	return func(c echo.Context) error {
		users, err := svc.GetAllUser(c.Request().Context())
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		// fmt.Println(users)
		return c.JSON(http.StatusOK, users)
	}
}

type GetUserByIdFn func(context.Context, string) (*Users, error)

func (gn GetUserByIdFn) GetUserById(ctx context.Context, str string) (*Users, error) {
	return gn(ctx, str)
}

func GetUserByIdHandler(svc GetUserByIdFn) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Param("Name")
		users, err := svc.GetUserById(c.Request().Context(), user)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, users)
	}
}
