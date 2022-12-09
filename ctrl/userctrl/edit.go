package userctrl

import (
	"fmt"
	"net/http"
	"studentHelper/svc/usersvc"

	"github.com/labstack/echo/v4"
)

func (c *ctrl) Edit(ctx echo.Context) error {
	payload := new(usersvc.EditRequest)
	if err := ctx.Bind(&payload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Errorf("%w: %s", http.ErrBodyNotAllowed, err))
	}

	res, err := c.svc.Edit(ctx.Request().Context(), payload)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, &res)
}
