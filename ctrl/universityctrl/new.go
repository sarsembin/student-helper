package universityctrl

import (
	"github.com/labstack/echo/v4"
	"studentHelper/svc/universitysvc"
)

type Ctrl interface {
	Get(echo.Context) error
	Post(echo.Context) error
	Put(echo.Context) error
	Delete(echo.Context) error
}

type ctrl struct {
	svc universitysvc.Service
}

func New(svc universitysvc.Service) Ctrl {
	return &ctrl{
		svc: svc,
	}
}
