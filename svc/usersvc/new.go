package usersvc

import (
	"context"
	"studentHelper/repo/postgredb/userdb"
)

type Service interface {
	Get(context.Context, *GetRequest) (*GetResponse, error)
	Edit(ctx context.Context, req *EditRequest) (*EditResponse, error)
	Delete(ctx context.Context, req *DeleteRequest) (*DeleteResponse, error)
	Register(ctx context.Context, req *RegisterRequest) (*RegisterResponse, error)
	Login(ctx context.Context, req *LoginRequest) (*LoginResponse, error)
}

type service struct {
	repo userdb.Repository
}

func New(repo userdb.Repository) Service {
	return &service{
		repo: repo,
	}
}
