package svc

import (
	"studentHelper/repo"
	"studentHelper/svc/universitysvc"
)

type Services struct {
	Unisvc universitysvc.Service
}

func Init(repos *repo.Repositories) *Services {
	return &Services{
		Unisvc: universitysvc.New(repos.Postgrerepo),
	}
}
