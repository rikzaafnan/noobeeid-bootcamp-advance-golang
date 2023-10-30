package product

import (
	"context"
	"errors"
	"github.com/jmoiron/sqlx"
	"log"
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
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func (p PostgresSQLXRepository) FindOneByProductID(_ context.Context, productID int64) (product Product, err error) {

	query := `
		SELECT id, name, category, price, stock FROM products WHERE id = $1
	`

	err = p.db.Get(&product, query, productID)
	if err != nil {
		log.Println(err)
		return
	}

	return product, nil
}

func (p PostgresSQLXRepository) UpdateProductByProductID(_ context.Context, productID int64, req Product) (err error) {

	req.Id = int(productID)

	row, err := p.db.NamedExec(`UPDATE products
				SET name = :name, category = :category, price = :price, stock = :stock
				WHERE id = :id`, req)
	if err != nil {
		log.Println(err)
		return
	}

	rowAffected, _ := row.RowsAffected()
	if rowAffected < 0 {
		return errors.New("internal server error")
	}

	return nil
}

func (p PostgresSQLXRepository) DeleteOneProductByProductID(_ context.Context, productID int64) (err error) {

	row := p.db.MustExec("DELETE FROM products WHERE id = $1", productID)

	rowAffected, _ := row.RowsAffected()
	if rowAffected < 0 {
		return errors.New("internal server error")
	}

	return nil
}
