package commentssvc

import (
	"context"
	"studentHelper/repo/postgredb/commentsdb"
)

type CreateRequest struct {
	JsonEntity
}

type CreateResponse struct {
}

func (s *service) Add(ctx context.Context, req *CreateRequest) (*CreateResponse, error) {
	err := s.repo.Add(&commentsdb.Entity{
		Content:      req.Content,
		UserID:       req.UserID,
		UniversityID: req.UniversityID,
	})
	if err != nil {
		return nil, err
	}

	return &CreateResponse{}, nil
}
