package userctrl

import (
	"fmt"
	"net/http"
	"studentHelper/svc/usersvc"

	"github.com/labstack/echo/v4"
)

func (c *ctrl) Login(ctx echo.Context) error {
	payload := new(usersvc.LoginRequest)
	if err := ctx.Bind(&payload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Errorf("%w: %s", http.ErrBodyNotAllowed, err))
	}

	res, err := c.svc.Login(ctx.Request().Context(), payload)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, &res)
}
