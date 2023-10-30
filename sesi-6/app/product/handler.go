package product

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strconv"
	"strings"
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

func (h Handler) FindAllProduct(c *fiber.Ctx) error {

	products, err := h.service.FindAllProduct(c.UserContext())
	if err != nil {
		payload := fiber.Map{}
		httpCode := 400
		payload = fiber.Map{
			"success": false,
			"message": "ERR INTERNAL",
			"error":   "ada masalah pada server",
		}

		httpCode = http.StatusInternalServerError
		return c.Status(httpCode).JSON(payload)
	}

	if len(products) <= 0 {
		payload := fiber.Map{}
		httpCode := 404
		payload = fiber.Map{
			"success": false,
			"message": "ERR NOT FOUND",
			"error":   "data tidak ditemukan",
		}

		httpCode = http.StatusNotFound
		return c.Status(httpCode).JSON(payload)
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "GET ALL SUCCESS",
		"payload": products,
	})
}

func (h Handler) FindOneProductByProductID(c *fiber.Ctx) error {

	productID := c.Params("productID")

	productIDInt, _ := strconv.Atoi(productID)

	product, err := h.service.FindOneProductByProductID(c.UserContext(), int64(productIDInt))
	if err != nil {
		payload := fiber.Map{}
		httpCode := 400
		payload = fiber.Map{
			"success": false,
			"message": "ERR INTERNAL",
			"error":   err,
		}

		httpCode = http.StatusInternalServerError

		stringError := fmt.Sprintf("%s", err)
		stringErr := strings.Contains(stringError, "record not found")

		if stringErr == true {
			payload = fiber.Map{
				"success": false,
				"message": "ERR NOT FOUND",
				"error":   stringError,
			}

			httpCode = http.StatusNotFound
			return c.Status(httpCode).JSON(payload)
		}
		return c.Status(httpCode).JSON(payload)

	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "GET DATA SUCCESS",
		"payload": product,
	})
}

func (h Handler) UpdateProductByProductID(c *fiber.Ctx) error {

	productID := c.Params("productID")

	productIDInt, err := strconv.Atoi(productID)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "ERR BAD REQUEST",
			"error":   err.Error(),
		})
	}

	var req CreateProductrequest

	err = c.BodyParser(&req)
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

	err = h.service.UpdateProductByProductID(c.UserContext(), int64(productIDInt), model)
	if err != nil {
		payload := fiber.Map{}
		httpCode := 400
		payload = fiber.Map{
			"success": false,
			"message": "ERR INTERNAL",
			"error":   err,
		}

		httpCode = http.StatusInternalServerError

		stringError := fmt.Sprintf("%s", err)
		stringErr := strings.Contains(stringError, "record not found")

		if stringErr == true {
			payload = fiber.Map{
				"success": false,
				"message": "ERR NOT FOUND",
				"error":   stringError,
			}

			httpCode = http.StatusNotFound
			return c.Status(httpCode).JSON(payload)
		}
		return c.Status(httpCode).JSON(payload)

	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "UPDATE SUCCESS",
	})
}

func (h Handler) DeleteProductByProductID(c *fiber.Ctx) error {

	productID := c.Params("productID")

	productIDInt, err := strconv.Atoi(productID)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "ERR BAD REQUEST",
			"error":   err.Error(),
		})
	}

	err = h.service.DeleteProductByProductID(c.UserContext(), int64(productIDInt))
	if err != nil {
		payload := fiber.Map{}
		httpCode := 400
		payload = fiber.Map{
			"success": false,
			"message": "ERR INTERNAL",
			"error":   err,
		}

		httpCode = http.StatusInternalServerError

		stringError := fmt.Sprintf("%s", err)
		stringErr := strings.Contains(stringError, "record not found")

		if stringErr == true {
			payload = fiber.Map{
				"success": false,
				"message": "ERR NOT FOUND",
				"error":   stringError,
			}

			httpCode = http.StatusNotFound
			return c.Status(httpCode).JSON(payload)
		}
		return c.Status(httpCode).JSON(payload)

	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "DELETE DATA SUCCESS",
	})
}
