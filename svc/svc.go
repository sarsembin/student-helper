package svc

import (
	"studentHelper/repo"
	"studentHelper/svc/universitescoresvc"
	"studentHelper/svc/universitysvc"
)

type Services struct {
	Unisvc      universitysvc.Service
	UniScoresvc universitescoresvc.Service
}

func Init(repos *repo.Repositories) *Services {
	return &Services{
		Unisvc: universitysvc.New(repos.Postgrerepo.Universitydb),
		UniScoresvc: universitescoresvc.New(repos.Postgrerepo.Universitescoredb),
	}
}
