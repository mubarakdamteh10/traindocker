package customer_order

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type updateOrderByParamHandlerFn func(context.Context, primitive.ObjectID, Order) error

func (fn updateOrderByParamHandlerFn) UpdateOrderByParam(ctx context.Context, objectId primitive.ObjectID, order Order) error {
	return fn(ctx, objectId, order)
}

func UpdateOrderByparamHandler(svc updateOrderByParamHandlerFn) echo.HandlerFunc {
	return func(c echo.Context) error {
		order := Order{}
		orderId := c.Param("ID")
		objectId, err := primitive.ObjectIDFromHex(orderId)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		if err := c.Bind(&order); err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		err = svc.UpdateOrderByParam(c.Request().Context(), objectId, order)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, "edited")
	}
}

// order update by field

type updateOrderByFieldFn func(context.Context, Order) error

func (fn updateOrderByFieldFn) UpdateOrderByField(ctx context.Context, order Order) error {
	return fn(ctx, order)
}

func UpdateOrderByFieldHandler(svc updateOrderByFieldFn) echo.HandlerFunc {
	return func(c echo.Context) error {
		order := Order{}
		if err := c.Bind(&order); err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		if err := svc.UpdateOrderByField(c.Request().Context(), order); err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, "edited field")
	}
}
