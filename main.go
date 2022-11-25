package main

import (
	"studentHelper/ctrl"
	"studentHelper/svc"

	"github.com/labstack/echo/v4"
)

func main() {
	// Echo instance
	e := echo.New()

	svc := svc.Init()

	ctrl := ctrl.Init(svc)

	ctrl.RoutesRegister(e)

	// Start server
	e.Logger.Fatal(e.Start(":9000"))
}
