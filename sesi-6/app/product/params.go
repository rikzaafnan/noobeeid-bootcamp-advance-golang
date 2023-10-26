package product

type CreateProductrequest struct {
	Name     string `json:"name"`
	Category string `json:"category"`
	Price    int    `json:"price"`
	Stock    int    `json:"stock"`
}
