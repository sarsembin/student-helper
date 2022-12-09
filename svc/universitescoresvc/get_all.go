package universitescoresvc

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
				Title: res[i].Title,
				Eval:  res[i].Eval,
				UniID: res[i].University_id,
			},
		})
	}

	return &GetAllResponse{Data: data}, nil
}
