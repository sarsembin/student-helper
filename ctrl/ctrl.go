package ctrl

import (
	universityctrl "studentHelper/ctrl/universityctrl"
	"studentHelper/svc"
)

type Controllers struct {
	unictrl universityctrl.Ctrl
}

func Init(svc *svc.Services) *Controllers {
	return &Controllers{
		unictrl: universityctrl.New(svc.Unisvc),
	}
}
