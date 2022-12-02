package universitysvc

import (
	"context"
	"studentHelper/repo/postgredb"
)

type Service interface {
	Add(context.Context, *CreateRequest) (*CreateResponse, error)
}

type service struct {
	postgredb postgredb.Repository
}

func New(pgdb postgredb.Repository) Service {
	return &service{
		postgredb: pgdb,
	}
}
