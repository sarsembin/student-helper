package ctrl

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (c *Controllers) RoutesRegister(e *echo.Echo) {
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.POST(prepend("/university"), c.unictrl.Post)
}

// 💩
func prepend(s string) string {
	return "/api/studentHelper" + s
}
