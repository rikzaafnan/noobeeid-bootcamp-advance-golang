package product

import (
	"context"
	"github.com/jmoiron/sqlx"
)

type PostgresSQLXRepository struct {
	db *sqlx.DB
}

func NewPostgresSQLXRepository(db *sqlx.DB) PostgresSQLXRepository {
	return PostgresSQLXRepository{
		db: db,
	}
}

func (p PostgresSQLXRepository) Create(ctx context.Context, model Product) (err error) {

	query := `
		INSERT INTO products ( name, category, price, stock) VALUES (:name, :category, :price, :stock)
	`

	stmt, err := p.db.PrepareNamed(query)
	if err != nil {
		return
	}

	defer stmt.Close()

	_, err = stmt.Exec(model)
	if err != nil {
		return
	}

	return
}
func (p PostgresSQLXRepository) FindAll(_ context.Context) (products []Product, err error) {

	query := `
		SELECT id, name, category, price, stock FROM products 
	`

	err = p.db.Select(&products, query)

	return
}

func (p PostgresSQLXRepository) FindOneByProductID(_ context.Context, productID int64) (product Product, err error) {

	//err = p.db.Where("id = ?", productID).First(&product).Error
	//if err != nil {
	//	return product, err
	//}

	return product, nil
}

func (p PostgresSQLXRepository) UpdateProductByProductID(_ context.Context, productID int64, req Product) (err error) {

	//err = p.db.Where("id = ?", productID).First(&product).Error
	//if err != nil {
	//	return product, err
	//}

	return nil
}

func (p PostgresSQLXRepository) DeleteOneProductByProductID(_ context.Context, productID int64) (err error) {

	//err = p.db.Where("id = ?", productID).First(&product).Error
	//if err != nil {
	//	return product, err
	//}

	return nil
}
