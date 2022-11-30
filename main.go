package main

import (
	"log"
	"studentHelper/ctrl"
	"studentHelper/repo"
	"studentHelper/res"
	"studentHelper/svc"

	"github.com/labstack/echo/v4"
)

func main() {
	// Echo instance
	e := echo.New()

	res, err := res.New()
	if err != nil {
		log.Fatalf("res init error: %s", err)
	}

	repo := repo.New(res)

	svc := svc.Init()

	ctrl := ctrl.Init(svc)

	ctrl.RoutesRegister(e)

	// Start server
	e.Logger.Fatal(e.Start(":9000"))
}
