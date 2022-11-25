package universitysvc

import (
	"context"
)

type Service interface {
	Add(context.Context, *CreateRequest) (*CreateResponse, error)
}

type service struct {
}

func New() Service {
	return &service{
	}
}