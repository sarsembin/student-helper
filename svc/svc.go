package svc

import (
	"studentHelper/repo"
	"studentHelper/svc/commentssvc"
	"studentHelper/svc/universitescoresvc"
	"studentHelper/svc/universitysvc"
	"studentHelper/svc/useracademicinfosvc"
	"studentHelper/svc/usersvc"
)

type Services struct {
	Unisvc              universitysvc.Service
	UniScoresvc         universitescoresvc.Service
	Usersvc             usersvc.Service
	UserAcademicInfosvc useracademicinfosvc.Service
	Commentssvc         commentssvc.Service
}

func Init(repos *repo.Repositories) *Services {
	return &Services{
		Unisvc:              universitysvc.New(repos.Postgrerepo.Universitydb),
		UniScoresvc:         universitescoresvc.New(repos.Postgrerepo.Universitescoredb),
		Usersvc:             usersvc.New(repos.Postgrerepo.Userdb),
		UserAcademicInfosvc: useracademicinfosvc.New(repos.Postgrerepo.UserAcademicInfodb),
		Commentssvc:         commentssvc.New(repos.Postgrerepo.Commentsdb),
	}
}
