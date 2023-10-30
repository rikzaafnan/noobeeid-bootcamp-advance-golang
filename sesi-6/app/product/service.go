package product

import (
	"context"
	"log"
)

type Repository interface {
	Create(ctx context.Context, req Product) (err error)
	FindAll(ctx context.Context) (products []Product, err error)
	FindOneByProductID(ctx context.Context, productID int64) (product Product, err error)
	UpdateProductByProductID(ctx context.Context, productID int64, req Product) (err error)
	DeleteOneProductByProductID(ctx context.Context, productID int64) (err error)
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

func (s Service) FindAllProduct(ctx context.Context) (products []Product, err error) {

	products, err = s.repo.FindAll(ctx)
	if err != nil {
		log.Println("error when try to create to database with error", err.Error())
		return nil, err
	}

	return products, nil
}

func (s Service) FindOneProductByProductID(ctx context.Context, productID int64) (product Product, err error) {

	product, err = s.repo.FindOneByProductID(ctx, productID)
	if err != nil {
		log.Println("error when try to get with error", err.Error())
		return product, err
	}

	return product, nil
}

func (s Service) UpdateProductByProductID(ctx context.Context, productID int64, req Product) (err error) {

	if err = req.Validate(); err != nil {
		log.Println("error when try to validate request with error", err.Error())
		return
	}

	product, err := s.repo.FindOneByProductID(ctx, productID)
	if err != nil {
		log.Println("error when try to get with error", err.Error())
		return err
	}

	product.Name = req.Name
	product.Category = req.Category
	product.Price = req.Price
	product.Stock = req.Stock

	if err = s.repo.UpdateProductByProductID(ctx, productID, product); err != nil {
		log.Println("error when try to create to database with error", err.Error())
		return
	}

	return nil
}

func (s Service) DeleteProductByProductID(ctx context.Context, productID int64) (err error) {

	_, err = s.repo.FindOneByProductID(ctx, productID)
	if err != nil {
		log.Println("error when try to get with error", err.Error())
		return err
	}

	if err = s.repo.DeleteOneProductByProductID(ctx, productID); err != nil {
		log.Println("error when try to delete to database with error", err.Error())
		return
	}

	return nil
}
