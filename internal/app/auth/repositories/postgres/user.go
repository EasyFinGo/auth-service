package postgres

import "github.com/jackc/pgx/v5/pgxpool"

type RegistrationRepository struct {
	queries *db.Queries
	pool    *pgxpool.Pool
}