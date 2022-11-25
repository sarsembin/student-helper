package svc

import "studentHelper/svc/universitysvc"

type Services struct {
	Unisvc universitysvc.Service
}

func Init() *Services {
	return &Services{
		Unisvc: universitysvc.New(),
	}
}
