package domain

import "errors"

var (
	ErrCategoryNotFound        = errors.New("category not found")
	ErrCategoryAlreadyExists   = errors.New("category already exists")
	ErrUnexpectedDatabase      = errors.New("unexpected database error")
	ErrCategoryHasTransactions = errors.New("category has transactions")
	ErrInvalidCategoryID       = errors.New("invalid category ID")
	ErrEntityDecoding          = errors.New("error decoding entity from database")
)
