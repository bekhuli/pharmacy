package user

import "errors"

var (
	ErrPhoneExists        = errors.New("phone is already exists")
	ErrInvalidCredentials = errors.New("invalid phone or password")
)
