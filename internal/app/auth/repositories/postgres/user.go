package postgres

import (
	"EasyFinGo/internal/app/auth/models"
	"context"
	"database/sql"
	"errors"
)

var (
	ErrPESELAlreadyExists = errors.New("pesel already registered")
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(ctx context.Context, user *models.User) (int64, error) {
	query := `
		INSERT INTO users (
			first_name, middle_name, last_name, date_of_birth, pesel, is_accepted
		) VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id
	`

	var id int64
	err := r.db.QueryRowContext(ctx, query,
		user.FirstName,
		user.MiddleName,
		user.LastName,
		user.DateOfBirth,
		user.Pesel,
		user.IsAccepted,
	).Scan(&id)

	if err != nil {
		return 0, err
	}

	user.ID = id
	return id, nil
}

func (r *userRepository) CreateAddress(ctx context.Context, addr *models.Address) error {
	query := `
		INSERT INTO user_addresses (
			user_id, postcode, street, building_no, apartment_no, city, country
		) VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	_, err := r.db.ExecContext(ctx, query,
		addr.UserID,
		addr.Postcode,
		addr.Street,
		addr.BuildingNo,
		addr.ApartmentNo,
		addr.City,
		addr.Country,
	)
	return err
}

func (r *userRepository) GetUserByPESEL(ctx context.Context, pesel string) (*models.User, error) {
	user := &models.User{}
	query := `
		SELECT id, first_name, middle_name, last_name, date_of_birth, pesel, is_accepted, created_at, updated_at
		FROM users
		WHERE pesel = $1
	`

	err := r.db.QueryRowContext(ctx, query, pesel).Scan(
		&user.ID,
		&user.FirstName,
		&user.MiddleName,
		&user.LastName,
		&user.DateOfBirth,
		&user.Pesel,
		&user.IsAccepted,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil 
	}
	if err != nil {
		return nil, err
	}
	return user, nil
}
