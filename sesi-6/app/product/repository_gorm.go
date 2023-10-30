package product

import (
	"context"
	"gorm.io/gorm"
)

type PostgresGormRepository struct {
	db *gorm.DB
}

func NewPostgresGormRepository(db *gorm.DB) PostgresGormRepository {
	return PostgresGormRepository{
		db: db,
	}
}

func (p PostgresGormRepository) Create(_ context.Context, model Product) (err error) {
	return p.db.Create(&model).Error
}

func (p PostgresGormRepository) FindAll(_ context.Context) (products []Product, err error) {

	err = p.db.Find(&products).Error
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (p PostgresGormRepository) FindOneByProductID(_ context.Context, productID int64) (product Product, err error) {

	err = p.db.Where("id = ?", productID).First(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil
}

func (p PostgresGormRepository) UpdateProductByProductID(_ context.Context, productID int64, req Product) (err error) {

	err = p.db.Where("id =  ?", productID).Save(&req).Error
	if err != nil {
		return err
	}

	return nil
}

func (p PostgresGormRepository) DeleteOneProductByProductID(_ context.Context, productID int64) (err error) {

	var product Product

	err = p.db.Where("id = ?", productID).Delete(&product).Error
	if err != nil {
		return err
	}

	return nil
}
