package product

import (
	"context"
	"log"
)

type Repository interface {
	Create(ctx context.Context, req Product) (err error)
	FindAll(ctx context.Context) (products []Product, err error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return Service{
		repo: repo,
	}
}

func (s Service) CreateProduct(ctx context.Context, req Product) (err error) {

	if err = req.Validate(); err != nil {
		log.Println("error when try to validate request with error", err.Error())
		return
	}

	if err = s.repo.Create(ctx, req); err != nil {
		log.Println("error when try to create to database with error", err.Error())
		return
	}

	return nil
}

func (s Service) FindAllProduct(ctx context.Context) (err error) {

	products, err := s.repo.FindAll(ctx)
	if err != nil {
		log.Println("error when try to create to database with error", err.Error())
		return
	}

	log.Println(products)

	return nil
}
