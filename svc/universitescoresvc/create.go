package universitescoresvc

import (
	"context"
	"studentHelper/repo/postgredb/universitescoredb"

)

type CreateRequest struct {
	JsonEntity
}

type CreateResponse struct {
}

func (s *service) Add(ctx context.Context, req *CreateRequest) (*CreateResponse, error) {
	err := s.repo.Add(&universitescoredb.Entity{
		Title: req.Title,
		Eval:  req.Eval,
		University_id: req.UniID,
	})
	if err != nil {
		return nil, err
	}

	return &CreateResponse{}, nil
}
