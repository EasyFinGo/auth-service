package err

import "errors"

var (
	ErrUserAlreadyExists = errors.New("user with this PESEL already exists")
	ErrUserNotFound      = errors.New("user not found")
	ErrDBInternal        = errors.New("internal database error")
)
