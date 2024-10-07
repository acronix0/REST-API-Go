package service

import (
	"context"

	"github.com/acronix0/REST-API-Go/internal/domain"
	"github.com/acronix0/REST-API-Go/internal/repository"
)
type categoriesService struct{
	repo repository.Category
}
func NewCategoriesService(repo repository.Category) *categoriesService{
	return &categoriesService{repo: repo}
}

func (s *categoriesService) GetCategories(ctx context.Context) ([]domain.Category, error){
  return s.repo.GetCategories(ctx)
}

