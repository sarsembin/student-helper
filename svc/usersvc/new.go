package usersvc

import (
	"context"
	"studentHelper/repo/postgredb/userdb"
)

type Service interface {
	Add(context.Context, *CreateRequest) (*CreateResponse, error)
	Get(context.Context, *GetRequest) (*GetResponse, error)
	GetAll(ctx context.Context, req *GetAllRequest) (*GetAllResponse, error)
	Put(ctx context.Context, req *PutRequest) (*PutResponse, error)
	Delete(ctx context.Context, req *DeleteRequest) (*DeleteResponse, error)
}

type service struct {
	repo userdb.Repository
}

func New(repo userdb.Repository) Service {
	return &service{
		repo: repo,
	}
}
