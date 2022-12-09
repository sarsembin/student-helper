package usersvc

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type jwtCustomClaims struct {
	Admin    bool   `json:"admin"`
	Username string `json:"username"`
	jwt.StandardClaims
}

const expirationTime = time.Hour

var ErrInvalidCredentials error = fmt.Errorf("invalid credentials")

func (s *service) Login(ctx context.Context, req *LoginRequest) (*LoginResponse, error) {
	acc, err := s.repo.GetByEmail(req.Email)
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	err = bcrypt.CompareHashAndPassword([]byte(acc.Password), []byte(req.Password))
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	claims := &jwtCustomClaims{
		acc.IsSuperuser,
		acc.Username,
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

	acc.LastLogin = time.Now()

	err = s.repo.Put(acc)
	if err != nil {
		return nil, err
	}

	return &LoginResponse{Token: t}, nil
}
