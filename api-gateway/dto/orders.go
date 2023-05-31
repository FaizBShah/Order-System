package dto

type Cart struct {
	Products []Product `json:"products"`
}

type CreateOrderRequest struct {
	Products []Product `json:"products"`
}

type CreateOrderResponse struct {
	Id     int64 `json:"id"`
	UserId int64 `json:"userId"`
	Cart   Cart  `json:"cart"`
}
