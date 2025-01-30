package onlineStore

type Seller struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Seller      Seller  `json:"seller"`
	SellerID    int     `json:"sellerId"`
}

type Customer struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

type Order struct {
	ID         int       `json:"id"`
	Customer   Customer  `json:"customer"`
	CustomerID int       `json:"customerId"`
	Product    []Product `json:"product"`
}

type UpdateSellerRequest struct {
	Name  *string `json:"name"`
	Phone *string `json:"phone"`
}
