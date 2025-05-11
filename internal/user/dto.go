package user

import "time"

type RegisterRequest struct {
	Phone     string `json:"phone" validate:"required,phone"`
	FirstName string `json:"first_name"  validate:"required"`
	LastName  string `json:"last_name"  validate:"required"`
	Password  string `json:"password"  validate:"required,min=8"`
}

type LoginRequest struct {
	Phone    string `json:"phone" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UpdateProfileRequest struct {
	FirstName *string `json:"first_name,omitempty"`
	LastName  *string `json:"last_name,omitempty"`
	Age       *int    `json:"age,omitempty"`
	Job       *string `json:"job,omitempty"`
	Gender    *string `json:"gender,omitempty"`
	IsMarried *bool   `json:"is_married,omitempty"`
}

type Response struct {
	ID        string    `json:"id"`
	Phone     string    `json:"phone"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	CreatedAt time.Time `json:"created_at"`
}

type PublicResponse struct {
	Phone     string  `json:"phone"`
	FirstName *string `json:"first_name"`
	LastName  *string `json:"last_name"`
	Age       *int    `json:"age"`
	Job       *string `json:"job"`
	Gender    *string `json:"gender"`
	IsMarried *bool   `json:"is_married"`
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

func ToPublicResponse(p *Profile) *PublicResponse {
	return &PublicResponse{
		Phone:     p.Phone,
		FirstName: p.FirstName,
		LastName:  p.LastName,
		Age:       p.Age,
		Job:       p.Job,
		Gender:    p.Gender,
		IsMarried: p.IsMarried,
	}
}
