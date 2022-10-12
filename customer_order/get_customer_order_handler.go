package customer_order

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
)

type getAllOrdersFn func(context.Context) ([]*Order, error)

func (gn getAllOrdersFn) GetAllOrder(ctx context.Context) ([]*Order, error) {
	return gn(ctx)
}

func GetAllOrderHandler(svc getAllOrdersFn) echo.HandlerFunc {
	return func(c echo.Context) error {
		orders, err := svc.GetAllOrder(c.Request().Context())
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, orders)
	}
}

type GetOrderByIdFn func(context.Context, string) (*Order, error)

func (gn GetOrderByIdFn) GetOrderById(ctx context.Context, str string) (*Order, error) {
	return gn(ctx, str)
}

func GetOrderByIdHandler(svc GetOrderByIdFn) echo.HandlerFunc {
	return func(c echo.Context) error {
		order := c.Param("Name")
		orders, err := svc.GetOrderById(c.Request().Context(), order)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, orders)
	}
}
