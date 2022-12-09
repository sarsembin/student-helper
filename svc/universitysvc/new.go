package universitysvc

import (
	"context"
	"studentHelper/repo/postgredb"
)

type Service interface {
	Add(context.Context, *CreateRequest) (*CreateResponse, error)
	Get(context.Context, *GetRequest) (*GetResponse, error)
	GetAll(ctx context.Context, req *GetAllRequest) (*GetAllResponse, error)
	Put(ctx context.Context, req *PutRequest) (*PutResponse, error)
	Delete(ctx context.Context, req *DeleteRequest) (*DeleteResponse, error)
}

type service struct {
	postgredb postgredb.Repository
}

func New(pgdb postgredb.Repository) Service {
	return &service{
		postgredb: pgdb,
	}
}
