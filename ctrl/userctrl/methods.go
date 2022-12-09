package userctrl

import (
	"net/http"
	"strconv"
	"studentHelper/svc/usersvc"

	"github.com/labstack/echo/v4"
)

// GET
func (c *ctrl) Get(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	res, err := c.svc.Get(ctx.Request().Context(), &usersvc.GetRequest{ID: id})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, &res)
}

// DELETE
func (c *ctrl) Delete(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	res, err := c.svc.Delete(ctx.Request().Context(), &usersvc.DeleteRequest{ID: id})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, &res)
}
