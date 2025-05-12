package admin

import (
	"time"

	"github.com/google/uuid"
)

type UsersResponse struct {
	ID        uuid.UUID `json:"id"`
	Phone     string    `json:"phone"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	CreatedAt time.Time `json:"created_at"`
}

func ToUsersResponse(u *User) *UsersResponse {
	return &UsersResponse{
		ID:        u.ID,
		Phone:     u.Phone,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		CreatedAt: u.CreatedAt,
	}
}
