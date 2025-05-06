package routes

import (
	"database/sql"
	"github.com/bekhuli/pharmacy/internal/user"

	"github.com/gorilla/mux"
)

func InitRouter(db *sql.DB) *mux.Router {
	r := mux.NewRouter()

	api := r.PathPrefix("/api/v1").Subrouter()

	userRepo := user.NewUserRepository(db)
	userValidator := user.NewUserValidator()
	userService := user.NewUserService(userRepo, userValidator)
	userHandler := user.NewUserHandler(userService)

	user.RegisterUserRoutes(api, userHandler)

	return r
}
