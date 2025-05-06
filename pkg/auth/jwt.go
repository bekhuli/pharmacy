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

func GenerateJWT(cfg common.JWTConfig, userID uuid.UUID, phone string) (string, error) {
	now := time.Now()
	expiration := now.Add(time.Duration(cfg.JWTExpiration) * time.Second)

	claims := CustomClaims{
		UserID: userID,
		Phone:  phone,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiration),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			Issuer:    "pharmacy-api",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(cfg.JWTSecret))
	if err != nil {
		return "", fmt.Errorf("JWT signing failed: %w", err)
	}

	return signedToken, nil
}
