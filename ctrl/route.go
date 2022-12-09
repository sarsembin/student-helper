package ctrl

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (c *Controllers) RoutesRegister(e *echo.Echo) {
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Group
	stHelper := e.Group("/api/studentHelper")
	// Routes
	stHelper.GET("/university/:id", c.unictrl.Get)
	stHelper.GET("/university", c.unictrl.GetAll)
	stHelper.POST("/university", c.unictrl.Post)
	stHelper.PUT("/university/:id", c.unictrl.Put)
	stHelper.DELETE("/university/:id", c.unictrl.Delete)

}
