package useracademicinfosvc

import (
	"context"
	"studentHelper/repo/postgredb/useracademicinfodb"
)

type CreateRequest struct {
	JsonEntity
}

type CreateResponse struct {
}

func (s *service) Add(ctx context.Context, req *CreateRequest) (*CreateResponse, error) {
	err := s.repo.Add(&useracademicinfodb.Entity{
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

	return &CreateResponse{}, nil
}
