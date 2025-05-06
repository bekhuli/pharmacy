package user

import "time"

type RegisterRequest struct {
	Phone     string `json:"phone" validate:"required"`
	FirstName string `json:"first_name"  validate:"required"`
	LastName  string `json:"last_name"  validate:"required"`
	Password  string `json:"password"  validate:"required"`
}

type Response struct {
	ID        string    `json:"id"`
	Phone     string    `json:"phone"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	CreatedAt time.Time `json:"created_at"`
}

func ToResponse(u *User) *Response {
	return &Response{
		ID:        u.ID.String(),
		Phone:     u.Phone,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		CreatedAt: u.CreatedAt,
	}
}
