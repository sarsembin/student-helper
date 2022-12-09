package useracademicinfosvc

import (
	"context"
	"studentHelper/repo/postgredb/useracademicinfodb"
)

type PutRequest struct {
	ID int `json:"id"`
	JsonEntity
}

type PutResponse struct {
}

func (s *service) Put(ctx context.Context, req *PutRequest) (*PutResponse, error) {
	err := s.repo.Put(&useracademicinfodb.Entity{
		AvgACTComposite: req.AvgACTComposite,
		AvgACTEnglish:   req.AvgACTEnglish,
		AvgACTMath:      req.AvgACTMath,
		AvgSATRW:        req.AvgSATRW,
		AvgSATMath:      req.AvgSATMath,
		UserID:          req.UserID,
	})
	if err != nil {
		return nil, err
	}

	return &PutResponse{}, nil
}
