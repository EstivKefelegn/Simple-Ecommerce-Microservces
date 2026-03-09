package postgres

type ProductModel struct {
	ID          string
	Name        string
	Description string
	Price       float64
	Stock       int64
}