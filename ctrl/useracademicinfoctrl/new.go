package useracademicinfoctrl

import (
	"studentHelper/svc/useracademicinfosvc"

	"github.com/labstack/echo/v4"
)

type Ctrl interface {
	Get(echo.Context) error
	GetAll(ctx echo.Context) error
	Post(echo.Context) error
	Put(echo.Context) error
	Delete(echo.Context) error
}

type ctrl struct {
	svc useracademicinfosvc.Service
}

func New(svc useracademicinfosvc.Service) Ctrl {
	return &ctrl{
		svc: svc,
	}
}
