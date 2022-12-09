package usersvc

import (
	"context"
	"net/mail"
	"studentHelper/repo/postgredb/userdb"
)

type EditRequest struct {
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Email     string `json:"email"`
}

type EditResponse struct {
}

func (s *service) Edit(ctx context.Context, req *EditRequest) (*EditResponse, error) {
	// sanitize email
	_, err := mail.ParseAddress(req.Email)
	if err != nil {
		return nil, err
	}

	newPass, err := SanitizeAndHash(req.Password)
	if err != nil {
		return nil, err
	}

	entity := &userdb.Entity{
		Password:  newPass,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Username:  req.Username,
		Email:     req.Email,
	}

	err = s.repo.Put(entity)
	if err != nil {
		return nil, err
	}

	return &EditResponse{}, nil
}
