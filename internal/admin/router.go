package admin

import "github.com/gorilla/mux"

func RegisterAdminRoutes(r *mux.Router, h *AdminHandler) {
	protected := r.PathPrefix("/admin").Subrouter()

	protected.HandleFunc("/users", h.GetAllUsers).Methods("GET")
}
