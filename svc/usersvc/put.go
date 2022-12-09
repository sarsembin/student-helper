package usersvc

import (
	"context"
	"studentHelper/repo/postgredb/userdb"
)

type PutRequest struct {
	ID int `json:"id"`
	JsonEntity
}

type PutResponse struct {
}

func (s *service) Put(ctx context.Context, req *PutRequest) (*PutResponse, error) {
	err := s.repo.Put(&userdb.Entity{
		ID:    req.ID,
		
	})
	if err != nil {
		return nil, err
	}

	return &PutResponse{}, nil
}
