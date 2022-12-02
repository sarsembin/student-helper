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

	res, err := res.Init()
	if err != nil {
		log.Fatalf("res init error: %s", err)
	}

	repo := repo.Init(res)

	svc := svc.Init(repo)

	ctrl := ctrl.Init(svc)

	ctrl.RoutesRegister(e)

	// Start server
	e.Logger.Fatal(e.Start(":9000"))
}
