package admin

import (
	"github.com/bekhuli/pharmacy/pkg/utils"
	"net/http"
)

type AdminHandler struct {
	service *Service
}

func NewAdminHandler(service *Service) *AdminHandler {
	return &AdminHandler{service: service}
}

func (h *AdminHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	users, err := h.service.GetAllUsers(ctx)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	var response []*UsersResponse
	for _, u := range users {
		response = append(response, ToUsersResponse(u))
	}

	utils.WriteJSON(w, http.StatusOK, response)
}
