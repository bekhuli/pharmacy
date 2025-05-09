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

type Profile struct {
	Phone     string
	FirstName string
	LastName  string
	Age       int64
	Job       string
	Gender    string
	IsMarried bool
}
