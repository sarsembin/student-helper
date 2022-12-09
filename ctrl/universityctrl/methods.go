package universityctrl

import (
	"fmt"
	"net/http"
	"strconv"
	"studentHelper/svc/universitysvc"

	"github.com/labstack/echo/v4"
)

// GET
func (c *ctrl) Get(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	res, err := c.svc.Get(ctx.Request().Context(), &universitysvc.GetRequest{ID: id})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, &res)
}

// GET ALL
func (c *ctrl) GetAll(ctx echo.Context) error {
	res, err := c.svc.GetAll(ctx.Request().Context(), &universitysvc.GetAllRequest{})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, &res)
}

// POST
func (c *ctrl) Post(ctx echo.Context) error {
	payload := new(universitysvc.CreateRequest)
	if err := ctx.Bind(&payload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Errorf("%w: %s", http.ErrBodyNotAllowed, err))
	}

	res, err := c.svc.Add(ctx.Request().Context(), payload)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, &res)
}

// PUT
func (c *ctrl) Put(ctx echo.Context) error {
	payload := new(universitysvc.PutRequest)
	if err := ctx.Bind(&payload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Errorf("%w: %s", http.ErrBodyNotAllowed, err))
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	payload.ID = id

	res, err := c.svc.Put(ctx.Request().Context(), payload)
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

	res, err := c.svc.Delete(ctx.Request().Context(), &universitysvc.DeleteRequest{ID: id})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, &res)
}
