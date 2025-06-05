package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const minKeySize = 32

type JWTMaker struct {
	secretKey string
}

func NewJWTMaker(secretKey string) (Maker, error) {
	if len(secretKey) < minKeySize {
		return nil, fmt.Errorf("invalid key size: must be at least %d characters", minKeySize)
	}

	return &JWTMaker{secretKey}, nil
}

func (maker *JWTMaker) CreateToken(username string, duration time.Duration) (string, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", err
	}

	jwtPayload := NewJWTPayloadClaims(payload)

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtPayload)

	token, err := jwtToken.SignedString(maker.secretKey)
	if err != nil {
		return "", err
	}

	return token, nil

}
func (maker *JWTMaker) VerifyToken(token string) (*Payload, error) {
	jwtToken, err := jwt.ParseWithClaims(token, &JWTPayloadClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidToken
		}
		return []byte(maker.secretKey), nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, ErrExpiredToken
		} else if errors.Is(err, ErrInvalidToken) {
			return nil, ErrInvalidToken
		} else {
			return nil, err
		}
	}

	claims, ok := jwtToken.Claims.(*JWTPayloadClaims)
	if !ok {
		return nil, ErrInvalidToken
	}

	return &claims.Payload, nil

}
