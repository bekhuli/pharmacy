package user

import (
	"net/http"

	"github.com/bekhuli/pharmacy/pkg/utils"
)

type UserHandler struct {
	service *Service
}

func NewUserHandler(service *Service) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var dto RegisterRequest
	if err := utils.BindJSON(r, &dto); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	user, err := h.service.RegisterUser(r.Context(), dto)
	if err != nil {
		utils.WriteError(w, http.StatusConflict, err)
	}

	utils.WriteJSON(w, http.StatusOK, ToResponse(user))
}
