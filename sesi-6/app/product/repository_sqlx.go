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

func (p PostgresSQLXRepository) CreateProductData(ctx context.Context, model Product) (err error) {

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
