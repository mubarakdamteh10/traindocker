package customer_order

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
)

type createOrderFn func(context.Context, *Order) error

func (fn createOrderFn) createOrderFn(ctx context.Context, order *Order) error {
	return fn(ctx, order)
}

func CreateOrderHandler(scu createOrderFn) echo.HandlerFunc {
	return func(c echo.Context) error {
		order := new(Order)
		if err := c.Bind(order); err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		if err := scu.createOrderFn(c.Request().Context(), order); err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, "created Order")
	}
}
