package svc

import (
	"studentHelper/repo"
	"studentHelper/svc/universitescoresvc"
	"studentHelper/svc/universitysvc"
	"studentHelper/svc/usersvc"
)

type Services struct {
	Unisvc      universitysvc.Service
	UniScoresvc universitescoresvc.Service
	Usersvc     usersvc.Service
}

func Init(repos *repo.Repositories) *Services {
	return &Services{
		Unisvc:      universitysvc.New(repos.Postgrerepo.Universitydb),
		UniScoresvc: universitescoresvc.New(repos.Postgrerepo.Universitescoredb),
		Usersvc:     usersvc.New(repos.Postgrerepo.Userdb),
	}
}
