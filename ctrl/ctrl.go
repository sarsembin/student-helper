package ctrl

import (
	"studentHelper/ctrl/universitescorectrl"
	"studentHelper/ctrl/universityctrl"
	"studentHelper/svc"
)

type Controllers struct {
	unictrl      universityctrl.Ctrl
	uniscorectrl universitescorectrl.Ctrl
}

func Init(svc *svc.Services) *Controllers {
	return &Controllers{
		unictrl:      universityctrl.New(svc.Unisvc),
		uniscorectrl: universitescorectrl.New(svc.UniScoresvc),
	}
}
