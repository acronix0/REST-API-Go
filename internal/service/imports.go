package service

import (
	"context"
	"errors"

	"github.com/acronix0/REST-API-Go/internal/domain"
	"github.com/acronix0/REST-API-Go/internal/repository"
)

type importService struct {
	categoryRepo repository.Category
	productRepo repository.Product
}

func NewImportsService(categoryRepo repository.Category,productRepo repository.Product) *importService {
	return &importService{categoryRepo: categoryRepo, productRepo: productRepo}
}
func (s *importService)ImportCategories(ctx context.Context, categories []domain.Category) error{
	return errors.New("notImplement")
}
func (s *importService)ImportProducts(ctx context.Context, products []domain.Product) error{
	return errors.New("notImplement")
}