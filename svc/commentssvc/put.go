package commentssvc

import (
	"context"
	"studentHelper/repo/postgredb/commentsdb"
)

type PutRequest struct {
	ID int `json:"id"`
	JsonEntity
}

type PutResponse struct {
}

func (s *service) Put(ctx context.Context, req *PutRequest) (*PutResponse, error) {
	err := s.repo.Put(&commentsdb.Entity{
		ID:           req.ID,
		Content:      req.Content,
		UserID:       req.UserID,
		UniversityID: req.UniversityID,
	})
	if err != nil {
		return nil, err
	}

	return &PutResponse{}, nil
}
