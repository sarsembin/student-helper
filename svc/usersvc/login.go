package usersvc

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json: "string"`
}


type jwtCustomClaims struct {
	Admin bool   `json:"admin"`
	jwt.StandardClaims
}

const expirationTime = time.Hour

func (s *service) Login(ctx context.Context, req *LoginRequest) (*LoginResponse, error) {
	acc, err := s.repo.GetByEmail(req.Email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(acc.Password), []byte(req.Password))
	if err != nil {
		return nil, err
	}

	claims := &jwtCustomClaims{
		acc.IsSuperuser,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(expirationTime).Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return nil, err
	}

	return &LoginResponse{Token: t}, nil
}
