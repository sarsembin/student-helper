package ctrl

import (
	"studentHelper/ctrl/universitescorectrl"
	"studentHelper/ctrl/universityctrl"
	"studentHelper/ctrl/useracademicinfoctrl"
	"studentHelper/ctrl/userctrl"
	"studentHelper/svc"
)

type Controllers struct {
	unictrl              universityctrl.Ctrl
	uniscorectrl         universitescorectrl.Ctrl
	userctrl             userctrl.Ctrl
	useracademicinfoctrl useracademicinfoctrl.Ctrl
}

func Init(svc *svc.Services) *Controllers {
	return &Controllers{
		unictrl:      universityctrl.New(svc.Unisvc),
		uniscorectrl: universitescorectrl.New(svc.UniScoresvc),
		userctrl:     userctrl.New(svc.Usersvc),
		useracademicinfoctrl: useracademicinfoctrl.New(svc.UserAcademicInfosvc),
	}
}
