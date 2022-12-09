package usersvc

import (
	"context"
	"fmt"
	"net/mail"
	"studentHelper/repo/postgredb/userdb"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type RegisterRequest struct {
	Password      string `json:"password"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	Username      string `json:"username"`
	Email         string `json:"email"`
	SuperUserSign string `json:"susign"`
}

type RegisterResponse struct {
}

const superusersign = "BermagambetDidNothingWrong!!!_727"

const minpasslen = 8

var ErrInvalidPass error = fmt.Errorf("invalid password")

func (s *service) Register(ctx context.Context, req *RegisterRequest) (*RegisterResponse, error) {
	// sanitize pass
	if len(req.Password) < minpasslen {
		return nil, fmt.Errorf("%w: %s", ErrInvalidPass, fmt.Sprintf("password is too short, should be at least %v long", minpasslen))
	}

	// sanitize email
	_, err := mail.ParseAddress(req.Email)
	if err != nil {
		return nil, err
	}

	// hash and salt, but dont bake yet
	pass, err := hashAndSalt([]byte(req.Password))
	if err != nil {
		return nil, err
	}

	// assign
	entity := userdb.Entity{
		Password:    pass,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		DateJoined:  time.Now(),
		Username:    req.Username,
		Email:       req.Email,
		IsActive:    true,
		IsSuperuser: false,
		IsStaff:     false,
	}

	// if reg is su
	if req.SuperUserSign == superusersign {
		entity.IsSuperuser = true
		entity.IsStaff = true
	}

	err = s.repo.Add(&entity)
	if err != nil {
		return nil, err
	}

	return &RegisterResponse{}, nil
}

func hashAndSalt(pwd []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}
