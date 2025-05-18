package admin

import (
	"time"

	"github.com/google/uuid"
)

type UserResponse struct {
	ID        uuid.UUID `json:"id"`
	Phone     string    `json:"phone"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}

type PaginationMetaDTO struct {
	Page       int `json:"page"`
	Limit      int `json:"limit"`
	Total      int `json:"total"`
	TotalPages int `json:"total_pages"`
}

type UserDTO struct {
	ID        uuid.UUID `json:"id"`
	Phone     string    `json:"phone"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}

type PaginatedResponseDTO[T any] struct {
	Data []T
	Meta PaginationMetaDTO
}

func ToUserResponse(u *User) *UserResponse {
	return &UserResponse{
		ID:        u.ID,
		Phone:     u.Phone,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		CreatedAt: u.CreatedAt,
		Role:      u.Role,
	}
}

func MapUsersToDTO(users []*User) []UserDTO {
	dtos := make([]UserDTO, 0, len(users))
	for _, u := range users {
		dtos = append(dtos, UserDTO{
			ID:        u.ID,
			Phone:     u.Phone,
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Role:      u.Role,
			CreatedAt: u.CreatedAt,
		})
	}

	return dtos
}
