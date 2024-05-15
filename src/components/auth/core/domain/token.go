package domain

import (
	"github.com/golang-jwt/jwt/v5"
	"gym-management/src/lib/primitives/application_specific"
	"gym-management/src/lib/primitives/generic"
	"time"
)

type Token string

type TokenClaims struct {
	UserId string
	Role   string
}

func NewToken(claims TokenClaims, secret string) (Token, *application_specific.ApplicationException) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": claims.UserId,
		"role":   claims.Role,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
		"iss":    "gym-management",
		"sub":    "auth",
		"aud":    claims.UserId,
		"iat":    time.Now().Unix(),
		"jti":    generic.GenerateULID(),
	})

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", application_specific.NewUnknownException("AUTH.TOKEN.CREATION_FAILED", err.Error(), nil)
	}

	return Token(tokenString), nil
}

func DecodeToken(tokenString string, secret string) (*TokenClaims, *application_specific.ApplicationException) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, application_specific.NewAuthenticationException("AUTH.TOKEN.INVALID", "Invalid token", map[string]string{
			"token": tokenString,
		})
	}

	if !token.Valid {
		return nil, application_specific.NewAuthenticationException("AUTH.TOKEN.INVALID", "Invalid token", map[string]string{
			"token": tokenString,
		})
	}

	storedClaims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, application_specific.NewAuthenticationException("AUTH.TOKEN.INVALID", "Invalid token", map[string]string{
			"token": tokenString,
		})
	}

	claims := TokenClaims{
		UserId: storedClaims["userId"].(string),
		Role:   storedClaims["role"].(string),
	}

	return &claims, nil
}

func (t Token) Value() string {
	return string(t)
}
