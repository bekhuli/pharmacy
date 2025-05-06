package auth

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/bekhuli/pharmacy/pkg/utils"

	"github.com/bekhuli/pharmacy/internal/common"
)

type ContextKey string

const UserKey ContextKey = "userID"

func JWTMiddleware(cfg common.JWTConfig) func(handler http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := w.Header().Get("Authorization")
			if authHeader == "" {
				utils.WriteError(w, http.StatusUnauthorized, errors.New("authorization header required"))
				return
			}

			tokenString := strings.TrimPrefix(authHeader, "Bearer ")
			if tokenString == "" {
				utils.WriteError(w, http.StatusUnauthorized, errors.New("bearer token required"))
			}

			claims, err := ParseJWT(tokenString, cfg)
			if err != nil {
				utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("error token: %v", err))
				return
			}

			ctx := context.WithValue(r.Context(), UserKey, claims.UserID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
