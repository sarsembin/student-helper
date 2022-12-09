package useracademicinfosvc

import (
	"context"
)

type GetAllRequest struct {
}

type GetAllResponse struct {
	Data []Data `json:"data"`
}

func (s *service) GetAll(ctx context.Context, req *GetAllRequest) (*GetAllResponse, error) {
	res, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	data := make([]Data, len(res))

	for i := range res {
		data = append(data, Data{
			ID: res[i].ID,
			JsonEntity: JsonEntity{
				AvgACTComposite: res[i].AvgACTComposite,
				AvgACTEnglish:   res[i].AvgACTEnglish,
				AvgACTMath:      res[i].AvgACTMath,
				AvgSATRW:        res[i].AvgSATRW,
				AvgSATMath:      res[i].AvgSATMath,
				UserID:          res[i].UserID,
			},
		})
	}

	return &GetAllResponse{Data: data}, nil
}
