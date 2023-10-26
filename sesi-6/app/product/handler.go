package product

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type Handler struct {
	service Service
}

func NewHandler(svc Service) Handler {
	return Handler{
		service: svc,
	}
}

func (h Handler) CreateProduct(c *fiber.Ctx) error {

	var req CreateProductrequest

	err := c.BodyParser(&req)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "ERR BAD REQUEST",
			"error":   err.Error(),
		})
	}

	model := Product{
		Name:     req.Name,
		Category: req.Category,
		Price:    req.Price,
		Stock:    req.Stock,
	}

	err = h.service.CreateProduct(c.UserContext(), model)
	if err != nil {
		payload := fiber.Map{}
		httpCode := 400
		switch err {
		case ErrEmptyPrice, ErrEmptyCategory, ErrEmptyName, ErrEmptyStock:
			payload = fiber.Map{
				"success": false,
				"message": "ERR BAD REQUEST",
				"error":   err.Error(),
			}

			httpCode = http.StatusBadRequest

		default:

			payload = fiber.Map{
				"success": false,
				"message": "ERR INTERNAL",
				"error":   err.Error(),
			}

			httpCode = http.StatusInternalServerError
		}
		return c.Status(httpCode).JSON(payload)
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "success create product",
	})
}
