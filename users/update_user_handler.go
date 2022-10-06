package users

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type updateUserByIdParamHandlerFn func(context.Context, primitive.ObjectID, Users) error

func (fn updateUserByIdParamHandlerFn) UpdateUserByIdParam(ctx context.Context, objectId primitive.ObjectID, user Users) error {
	return fn(ctx, objectId, user)
}

func UpdateUserByIdParamHandler(svc updateUserByIdParamHandlerFn) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := Users{}
		userId := c.Param("ID")
		objectId, err := primitive.ObjectIDFromHex(userId)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		if err := c.Bind(&user); err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		err = svc.UpdateUserByIdParam(c.Request().Context(), objectId, user)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, "edited")
	}
}

// user update field
type updateUserByFieldIdHandlerFn func(context.Context, Users) error

func (fn updateUserByFieldIdHandlerFn) UpdateUserByIdField(ctx context.Context, user Users) error {
	return fn(ctx, user)
}

func UpdateUserByIdFieldHandler(svc updateUserByFieldIdHandlerFn) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := Users{}
		if err := c.Bind(&user); err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		if err := svc.UpdateUserByIdField(c.Request().Context(), user); err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, user)
	}
}
