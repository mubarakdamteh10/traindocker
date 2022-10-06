package users

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type deleteUserHandlerFn func(context.Context, primitive.ObjectID) error

func (fn deleteUserHandlerFn) DeleteUserById(ctx context.Context, objectId primitive.ObjectID) error {
	return fn(ctx, objectId)
}

func DeleteUserByIdHandler(svc deleteUserHandlerFn) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Param("ID")
		objectId, err := primitive.ObjectIDFromHex(user)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		err = svc.DeleteUserById(c.Request().Context(), objectId)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		return c.JSON(http.StatusOK, "deleted User")
	}
}
