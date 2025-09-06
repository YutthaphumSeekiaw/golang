package order

type Order struct {
	ID       int     `json:"id"`
	Product  string  `json:"product"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}
