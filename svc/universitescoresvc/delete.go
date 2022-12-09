package universitescoresvc

import (
	"context"
)

type DeleteRequest struct {
	ID int `json:"id"`
}

type DeleteResponse struct {
}

func (s *service) Delete(ctx context.Context, req *DeleteRequest) (*DeleteResponse, error) {
	err := s.repo.Delete(req.ID)
	if err != nil {
		return nil, err
	}

	return &DeleteResponse{}, nil
}
