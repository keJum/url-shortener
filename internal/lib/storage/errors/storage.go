package errors

import "errors"

var (
	ErrUrlNotFound = errors.New("url not found")
	ErrUrlExists   = errors.New("url already exists")
)
