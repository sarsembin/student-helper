package useracademicinfosvc

import (
	"context"
)

type GetRequest struct {
	ID int `json:"id"`
}

type GetResponse struct {
	Data
}

type Data struct {
	ID int `json:"id"`
	JsonEntity
}

func (s *service) Get(ctx context.Context, req *GetRequest) (*GetResponse, error) {
	res, err := s.repo.Get(req.ID)
	if err != nil {
		return nil, err
	}

	return &GetResponse{Data: Data{
		ID: req.ID,
		JsonEntity: JsonEntity{
			AvgACTComposite: res.AvgACTComposite,
			AvgACTEnglish:   res.AvgACTEnglish,
			AvgACTMath:      res.AvgACTMath,
			AvgSATRW:        res.AvgSATRW,
			AvgSATMath:      res.AvgSATMath,
			UserID:          res.UserID,
		},
	}}, nil
}
