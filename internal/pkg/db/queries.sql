-- queries.sql

-- name: CreateInitialRegistration :one
INSERT INTO users (
    pesel,
    first_name,
    middle_name,
    last_name,
    date_of_birth
) VALUES ($1, $2, $3, $4, $5)
RETURNING id, pesel, first_name, middle_name, last_name, date_of_birth, created_at, updated_at;

-- name: GetUserByPesel :one
SELECT id, pesel, first_name, middle_name, last_name, date_of_birth, created_at, updated_at
FROM users
WHERE pesel = $1;
