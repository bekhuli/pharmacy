package user

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID
	Phone     string
	FirstName string
	LastName  string
	Password  string
	CreatedAt time.Time
	IsDeleted bool
}
