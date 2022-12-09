package ctrl

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (c *Controllers) RoutesRegister(e *echo.Echo) {
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:  []byte("secret"),
		TokenLookup: "query:token",
	}))

	// Group
	stHelper := e.Group("/api/studentHelper")
	user := e.Group("/user")
	// University
	stHelper.GET("/university/:id", c.unictrl.Get)
	stHelper.GET("/university", c.unictrl.GetAll)
	stHelper.POST("/university", c.unictrl.Post)
	stHelper.PUT("/university/:id", c.unictrl.Put)
	stHelper.DELETE("/university/:id", c.unictrl.Delete)
	// Universite score
	stHelper.GET("/universiteScores/:id", c.uniscorectrl.Get)
	stHelper.GET("/universiteScores", c.uniscorectrl.GetAll)
	stHelper.POST("/universiteScores", c.uniscorectrl.Post)
	stHelper.PUT("/universiteScores/:id", c.uniscorectrl.Put)
	stHelper.DELETE("/universiteScores/:id", c.uniscorectrl.Delete)
	// User
	user.POST("/register", c.userctrl.Register)
	user.POST("/login", c.userctrl.Login)
	user.POST("/edit", c.userctrl.Edit)
	// User academic info
	user.GET("/academicInfo", c.uniscorectrl.Get)
	user.POST("/academicInfo", c.uniscorectrl.Post)
	user.PUT("/academicInfo", c.uniscorectrl.Put)
	user.DELETE("/academicInfo", c.uniscorectrl.Delete)
}
