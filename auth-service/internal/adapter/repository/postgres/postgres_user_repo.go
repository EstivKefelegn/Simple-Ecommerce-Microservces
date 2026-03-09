package postgres

import (
	"context"
	"github/ecommerceMSAuth/internal/domain"
	"github/ecommerceMSAuth/pkg"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresUserRepo struct {
	db *pgxpool.Pool
}

func NewPostgresUserRepo(db *pgxpool.Pool) *PostgresUserRepo {
	return &PostgresUserRepo{
		db: db,
	}
}

func (r *PostgresUserRepo) Save(user *domain.User) error {

	if err := domain.ValidateUser(user); err != nil {
		return pkg.ErrorHandler(err, "invalid user")
	}

	model := ToModel(user)

	hashPassword, err := pkg.HashPassword(model.Password)
	if err != nil {
		return pkg.ErrorHandler(err, "could not hash password")
	}

	query := `
	INSERT INTO users (id, username, email, password)
	VALUES ($1,$2,$3,$4)
	`

	_, err = r.db.Exec(
		context.Background(),
		query,
		model.ID,
		model.Username,
		model.Email,
		hashPassword,
	)

	if err != nil {
		return pkg.ErrorHandler(err, "failed to save user")
	}

	return nil
}

func (r *PostgresUserRepo) FindByEmail(email string) (*domain.User, error) {

	if err := domain.ValidateEmail(email); err != nil {
		return nil, pkg.ErrorHandler(err, "invalid email")
	}

	query := `
	SELECT id, username, email, password
	FROM users
	WHERE email=$1
	`

	var model User

	err := r.db.QueryRow(
		context.Background(),
		query,
		email,
	).Scan(
		&model.ID,
		&model.Username,
		&model.Email,
		&model.Password,
	)

	if err != nil {
		return nil, pkg.ErrorHandler(err, "user not found")
	}

	return ToDomain(&model), nil
}
