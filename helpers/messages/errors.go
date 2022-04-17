package messages

import "errors"

var (
	ErrInvalidData   = errors.New("data invalid")
	ErrInsertSuccess = errors.New("data inserted successfully")
)
