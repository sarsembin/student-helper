package usersvc

import (
	"context"
	"studentHelper/repo/postgredb/userdb"
)

type CreateRequest struct {
	JsonEntity
}

type CreateResponse struct {
}

func (s *service) Add(ctx context.Context, req *CreateRequest) (*CreateResponse, error) {
	err := s.repo.Add(&userdb.Entity{})
	if err != nil {
		return nil, err
	}

	return &CreateResponse{}, nil
}
