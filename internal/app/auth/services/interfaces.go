package services

import (
	"EasyFinGo/internal/app/auth/models"
	"context"
)

type RegistrationService interface {
	RegisterPersonalInfo(ctx context.Context, user *models.User) (int64, error)
	AddResidentialAddress(ctx context.Context, pesel string, addr *models.Address) error
}
