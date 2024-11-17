package service

import (
	"context"
	"encoding/json"

	"github.com/acronix0/REST-API-Go/internal/domain"
	"github.com/acronix0/REST-API-Go/internal/kafka"
	"github.com/acronix0/REST-API-Go/internal/repository"
)
type categoriesService struct{
	repo repository.Category
	producer *kafka.KafkaProducer
}
func NewCategoriesService(repo repository.Category, producer *kafka.KafkaProducer) *categoriesService{
	return &categoriesService{repo: repo, producer: producer}
}

func (s *categoriesService) GetCategories(ctx context.Context) ([]domain.Category, error){
	testOrder := domain.Order{}
	jsonData, _ := json.Marshal(testOrder)
	s.producer.SendMessage(kafka.Orderopic, []byte{1}, jsonData)
  return s.repo.GetCategories(ctx)
}

