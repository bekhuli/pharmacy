package admin

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID
	Phone     string
	FirstName string
	LastName  string
	CreatedAt time.Time
	Role      string
}
