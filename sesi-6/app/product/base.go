package product

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"gorm.io/gorm"
	"log"
)

func RegisterServiceProduct(router fiber.Router, db *gorm.DB, dbSQLX *sqlx.DB, connectionDBLib string) {

	var repoProduct Repository

	switch connectionDBLib {
	case "gorm":
		repoProduct = NewPostgresGormRepository(db)
		log.Println("db gorm is used")
	case "sqlx":
		repoProduct = NewPostgresSQLXRepository(dbSQLX)
		log.Println("db sqlx is used")
	}

	//repoProduct = NewPostgresGormRepository(db)
	svcProduct := NewService(repoProduct)
	handler := NewHandler(svcProduct)

	var productRouter = router.Group("/products")
	{
		productRouter.Post("", handler.CreateProduct)
	}

}
