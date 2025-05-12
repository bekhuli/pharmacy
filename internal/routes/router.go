package routes

import (
	"database/sql"

	"github.com/bekhuli/pharmacy/internal/admin"
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

	adminRepo := admin.NewAdminRepository(db)
	adminValidator := admin.NewAdminValidator()
	adminService := admin.NewAdminService(adminRepo, adminValidator)
	adminHandler := admin.NewAdminHandler(adminService)

	admin.RegisterAdminRoutes(api, adminHandler)

	return r
}
