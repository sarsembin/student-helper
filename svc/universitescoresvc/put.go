package universitescoresvc

import (
	"context"
	"studentHelper/repo/postgredb/universitescoredb"
)

type PutRequest struct {
	ID int `json:"id"`
	JsonEntity
}

type PutResponse struct {
}

func (s *service) Put(ctx context.Context, req *PutRequest) (*PutResponse, error) {
	err := s.repo.Put(&universitescoredb.Entity{
		ID:    req.ID,
		Title: req.Eval,
		Eval:  req.Eval,
	})
	if err != nil {
		return nil, err
	}

	return &PutResponse{}, nil
}
