package userctrl

import (
	"studentHelper/svc/usersvc"

	"github.com/labstack/echo/v4"
)

type Ctrl interface {
	Get(echo.Context) error
	Edit(ctx echo.Context) error
	Delete(echo.Context) error
	Register(ctx echo.Context) error
	Login(ctx echo.Context) error
}

type ctrl struct {
	svc usersvc.Service
}

func New(svc usersvc.Service) Ctrl {
	return &ctrl{
		svc: svc,
	}
}
