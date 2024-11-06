package service

import (
	"context"
	"encoding/xml"
	"errors"
	"mime/multipart"

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
func (s *importService)Parse(file1, file2 *multipart.File) error{
	decoder := xml.NewDecoder(*file1)
	for{
		tok, err := decoder.Token()
		if err != nil {
			break
		}
		switch elem := tok.(type) {
		case xml.StartElement:
			if elem.Name.Local == "Группы"{

			}
		}
	}
	return nil
}
func (s *importService)ImportCategories(ctx context.Context, categories []domain.Category) error{
	return errors.New("notImplement")
}
func (s *importService)ImportProducts(ctx context.Context, products []domain.Product) error{
	return errors.New("notImplement")
}
func (s *importService)ImportPicture(file *multipart.File) error{
	return errors.New("notImplement")
}