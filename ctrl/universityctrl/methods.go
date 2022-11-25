package universityctrl

import (
	"fmt"
	"net/http"
	"studentHelper/svc/universitysvc"

	"github.com/labstack/echo/v4"
)

func (c *ctrl) Get(ctx echo.Context) error {
	return nil
}

func (c *ctrl) Post(ctx echo.Context) error {
	payload := new(universitysvc.CreateRequest)
	if err := ctx.Bind(&payload); err != nil {
		return fmt.Errorf("%w: %s", http.ErrBodyNotAllowed, err)
	}

	res, err := c.svc.Add(ctx.Request().Context(), payload)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, &res)
}

func (c *ctrl) Put(ctx echo.Context) error {
	return nil
}

func (c *ctrl) Delete(ctx echo.Context) error {
	return nil
}