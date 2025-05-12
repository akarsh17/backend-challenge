package models

type Order struct {
	ID       string      `json:"id"`
	Items    []OrderItem `json:"items"`
	Products []Product   `json:"products"`
}

type OrderItem struct {
	ProductID int64 `json:"productId"`
	Quantity  int   `json:"quantity"`
}

type OrderRequest struct {
	CouponCode string      `json:"couponCode,omitempty"`
	Items      []OrderItem `json:"items"`
}
