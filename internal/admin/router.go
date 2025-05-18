package admin

import (
	"github.com/bekhuli/pharmacy/internal/common"
	"github.com/bekhuli/pharmacy/pkg/auth"

	"github.com/gorilla/mux"
)

func RegisterAdminRoutes(r *mux.Router, h *AdminHandler) {
	protected := r.PathPrefix("/admin").Subrouter()
	protected.Use(auth.JWTMiddleware(common.JWTEnv))
	protected.Use(auth.RequireRole("admin"))

	protected.HandleFunc("/users", h.GetAllUsers).Methods("GET")
	protected.HandleFunc("/user/{id}", h.GetUser).Methods("GET")
}
