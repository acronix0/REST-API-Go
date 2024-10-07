package service

import (
	"context"

	"github.com/acronix0/REST-API-Go/internal/domain"
	"github.com/acronix0/REST-API-Go/internal/repository"
)

type orderService struct{
	orderRepo repository.Order
}
func NewOrdersService(repo repository.Order) *orderService{
	return &orderService{orderRepo: repo}
}

func (s *orderService) Create(ctx context.Context, orderInput CreateOrderInput) error{
	
	return s.orderRepo.Create(
		ctx, 
		mapOrderInputToRepo(orderInput),
	)
}

func (s *orderService) GetByUserId(ctx context.Context, userID int) ([]domain.Order, error){
	return s.orderRepo.GetByUserId(ctx, userID)
}

func mapProductsInputToRepo(products []ProductInput) []repository.ProductInput {
	repoProducts := make([]repository.ProductInput, len(products))
	for i, product := range products {
		repoProducts[i] = mapProductInputToRepo(product)
	}
	return repoProducts
}

func mapProductInputToRepo(product ProductInput) repository.ProductInput {
	return repository.ProductInput{
		ID:       product.ID,
		Article:  product.Article,
		Quantity: product.Quantity,
		Price:    product.Price,
		Image:    product.Image,
	}
}


func mapOrderInputToRepo(order CreateOrderInput) repository.CreateOrderInput {
  return repository.CreateOrderInput{
    ID:             order.ID,
    UserID:         order.UserID,
    Products:       mapProductsInputToRepo(order.Products), 
    TotalPrice:     order.TotalPrice,
    DeliveryType:   order.DeliveryType,
    RecipientName:  order.RecipientName,
    RecipientPhone: order.RecipientPhone,
    RecipientEmail: order.RecipientEmail,
    Address:        order.Address,
    Comment:        order.Comment,
  }
}