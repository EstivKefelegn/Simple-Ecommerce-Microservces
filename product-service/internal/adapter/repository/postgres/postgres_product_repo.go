package postgres

import (
	"context"
	"github/productMCS/internal/domain"
	"github/productMCS/pkg"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ProductRepo struct {
	db *pgxpool.Pool
}

func NewProductRepo(db *pgxpool.Pool) *ProductRepo {
	return &ProductRepo{db: db}
}

func (r *ProductRepo) Save(ctx context.Context, product *domain.Product) error {

	query := `
	INSERT INTO products(id,name,description,price,stock)
	VALUES($1,$2,$3,$4,$5)
	`
	if product.ID == uuid.Nil {
		product.ID = pkg.GenerateUUID()
	}

	_, err := r.db.Exec(
		ctx,
		query,
		product.ID,
		product.Name,
		product.Description,
		product.Price,
		product.Stock,
	)

	return err
}

func (r *ProductRepo) FindByID(ctx context.Context, id uuid.UUID) (*domain.Product, error) {

	query := `
	SELECT id,name,description,price,stock
	FROM products
	WHERE id=$1
	`

	row := r.db.QueryRow(ctx, query, id)

	var product domain.Product

	err := row.Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.Price,
		&product.Stock,
	)

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *ProductRepo) FindAll(ctx context.Context) ([]*domain.Product, error) {

	query := `
	SELECT id,name,description,price,stock
	FROM products
	`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*domain.Product

	for rows.Next() {
		var p domain.Product

		err := rows.Scan(
			&p.ID,
			&p.Name,
			&p.Description,
			&p.Price,
			&p.Stock,
		)

		if err != nil {
			return nil, err
		}

		products = append(products, &p)
	}

	return products, nil
}
