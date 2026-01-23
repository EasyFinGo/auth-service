package postgres

import (
	"EasyFinGo/internal/app/auth/models"
	"context"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) (int64, error)
	CreateAddress(ctx context.Context, addr *models.Address) error
	GetUserByPESEL(ctx context.Context, pesel string) (*models.User, error)
}
