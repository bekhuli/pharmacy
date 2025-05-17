package auth

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/bekhuli/pharmacy/internal/common"
	"github.com/bekhuli/pharmacy/pkg/utils"

	"github.com/google/uuid"
)

type ContextKey string

const (
	UserKey ContextKey = "userID"
	RoleKey ContextKey = "role"
)

func JWTMiddleware(cfg common.JWTConfig) func(handler http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				utils.WriteError(w, http.StatusUnauthorized, errors.New("authorization header required"))
				return
			}

			tokenString := strings.TrimPrefix(authHeader, "Bearer ")
			if tokenString == "" {
				utils.WriteError(w, http.StatusUnauthorized, errors.New("bearer token required"))
				return
			}

			claims, err := ParseJWT(tokenString, cfg)
			if err != nil {
				utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("error token: %v", err))
				return
			}

			ctx := context.WithValue(r.Context(), UserKey, claims.UserID)
			ctx = context.WithValue(r.Context(), RoleKey, claims.Role)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func RequireRole(roleReq string) func(handler http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			role, ok := GetRoleFromContext(r.Context())
			if !ok {
				utils.WriteError(w, http.StatusForbidden, errors.New("role missing in context"))
				return
			}

			if role == roleReq {
				next.ServeHTTP(w, r)
				return
			}

			utils.WriteError(w, http.StatusForbidden, errors.New("insufficient permission"))
		})
	}
}

func GetUserIDFromContext(ctx context.Context) (uuid.UUID, bool) {
	userID, ok := ctx.Value(UserKey).(uuid.UUID)
	return userID, ok
}

func GetRoleFromContext(ctx context.Context) (string, bool) {
	role, ok := ctx.Value(RoleKey).(string)
	return role, ok
}
