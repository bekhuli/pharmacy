package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/bekhuli/pharmacy/internal/common"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func ParseJWT(tokenString string, cfg common.JWTConfig) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&CustomClaims{},
		func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}

			return []byte(cfg.JWTSecret), nil
		},
	)

	if err != nil {
		return nil, fmt.Errorf("token parsing error: %w", err)
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claim")
	}

	if time.Now().After(claims.ExpiresAt.Time) {
		return nil, errors.New("token expired")
	}

	if claims.UserID == uuid.Nil {
		return nil, errors.New("invalid user ID")
	}

	return claims, nil
}
