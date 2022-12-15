package ctrl

import (
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const SigningKey = "notasecret"

func getJwtConfig() middleware.JWTConfig {
	return middleware.JWTConfig{
		SigningKey:     []byte(SigningKey),
		TokenLookup:    "header:Authorization",
		ParseTokenFunc: parseToken,
	}
}

func parseToken(auth string, c echo.Context) (interface{}, error) {
	keyFunc := func(t *jwt.Token) (interface{}, error) {
		if t.Method.Alg() != "HS256" {
			return nil, fmt.Errorf("unexpected jwt signing method=%v", t.Header["alg"])
		}

		return SigningKey, nil
	}

	// claims are of type `jwt.MapClaims` when token is created with `jwt.Parse`
	token, err := jwt.Parse(auth, keyFunc)
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		fmt.Println(token.Raw)
		return nil, errors.New("invalid token")
	}

	fmt.Println(token.Claims)
	
	return token, nil
}
