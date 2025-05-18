package admin

import (
	"math"
	"net/http"
	"strconv"

	"github.com/bekhuli/pharmacy/pkg/utils"
)

type AdminHandler struct {
	service *Service
}

func NewAdminHandler(service *Service) *AdminHandler {
	return &AdminHandler{service: service}
}

func (h *AdminHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	page, limit := parsePaginationParams(r)

	users, total, err := h.service.GetAllUsers(r.Context(), page, limit)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	meta := PaginationMetaDTO{
		Page:       page,
		Limit:      limit,
		Total:      total,
		TotalPages: int(math.Ceil(float64(total) / float64(limit))),
	}

	response := PaginatedResponseDTO[UserDTO]{
		Data: MapUsersToDTO(users),
		Meta: meta,
	}

	utils.WriteJSON(w, http.StatusOK, response)
}

func parsePaginationParams(r *http.Request) (page, limit int) {
	query := r.URL.Query()

	page, err := strconv.Atoi(query.Get("page"))
	if err != nil || page < 1 {
		page = 1
	}

	limit, err = strconv.Atoi(query.Get("limit"))
	if err != nil || limit < 1 || limit > 100 {
		limit = 10
	}

	return page, limit
}
