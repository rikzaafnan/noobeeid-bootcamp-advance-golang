package product

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterServiceProduct(router fiber.Router, db *gorm.DB) {

	repoProduct := NewPostgresGormRepository(db)
	svcProduct := NewService(repoProduct)
	handler := NewHandler(svcProduct)

	var productRouter = router.Group("/products")
	{
		productRouter.Post("", handler.CreateProduct)
	}

}
