package universitescorectrl

import (
	"studentHelper/svc/universitescoresvc"

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
	svc universitescoresvc.Service
}

func New(svc universitescoresvc.Service) Ctrl {
	return &ctrl{
		svc: svc,
	}
}