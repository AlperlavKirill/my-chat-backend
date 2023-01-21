package errors

import "errors"

var (
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrUserNotExists     = errors.New("user does not exists")
	ErrWrongPassword     = errors.New("wrong password")
)
