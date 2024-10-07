package service

import (
	"context"

	"github.com/acronix0/REST-API-Go/internal/domain"
	"github.com/acronix0/REST-API-Go/internal/repository"
)

type productService struct {
	repo repository.Product
}

func NewProductsService(repo repository.Product) *productService {
  return &productService{repo: repo}
}
func (s *productService) GetProducts(ctx context.Context) ([]domain.Product, error){
	return s.repo.GetProducts(ctx)
}
func (s *productService) GetByCredentials(ctx context.Context, query domain.GetProductsQuery) ([]domain.Product, error){
	return s.repo.GetByCredentials(ctx, query)
}
