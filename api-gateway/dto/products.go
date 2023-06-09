package dto

type Product struct {
	Id          int32   `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Quantity    int32   `json:"quantity"`
}

type CreateProductRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Quantity    int32   `json:"quantity"`
}

type DeletedProductResponse struct {
	Id int32 `json:"id"`
}

type UpdateProductQuantityRequest struct {
	Id       int32 `json:"id"`
	Quantity int32 `json:"quantity"`
}
