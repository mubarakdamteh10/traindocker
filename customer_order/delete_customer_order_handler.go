package customer_order

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type deleteOrderHandlerFn func(context.Context, primitive.ObjectID) error

func (fn deleteOrderHandlerFn) DeleteOrderById(ctx context.Context, objectId primitive.ObjectID) error {
	return fn(ctx, objectId)
}
func DeleteOrderByIdHandler(svc deleteOrderHandlerFn) echo.HandlerFunc {
	return func(c echo.Context) error {
		order := c.Param("ID")
		objectId, err := primitive.ObjectIDFromHex(order)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		err = svc.DeleteOrderById(c.Request().Context(), objectId)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		return c.JSON(http.StatusOK, "deleted order")
	}
}
