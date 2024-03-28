package storage

import "errors"

var (
	ErrNotFound = errors.New("not found")
	ErrAlreadyExists = errors.New("already exist")
	CodeNotFound = "no rows in result set"
	CodeAlreadyExists = "23505"
)
