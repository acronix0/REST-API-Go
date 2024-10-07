package domain

import "errors"

var (
	ErrCategoryNotFound = errors.New("category not found")
	ErrProductNotFound  = errors.New("product not found")
	ErrOrderNotFound    = errors.New("order not found")
	ErrUserAlredyExist  = errors.New("user with such email already exist")
	ErrUserNotFound     = errors.New("user with such email not found ")
	ErrRTokenNotFound   = errors.New("refresh token not found")
)
