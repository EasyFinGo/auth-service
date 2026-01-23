package services

import (
	"EasyFinGo/internal/app/auth/models"
	"EasyFinGo/internal/app/auth/repositories/postgres"
	"EasyFinGo/internal/app/auth/validator"
	"context"
	"errors"
)

var (
	ErrUserAlreadyRegistered = errors.New("user with this PESEL is already registered")
	ErrInvalidData           = errors.New("invalid registration data")
	ErrUserNotFound          = errors.New("user not found")
	ErrAddressAlreadyExists  = errors.New("user already has a residential address")
	ErrInvalidAddressData    = errors.New("invalid address data")
	ErrDatabaseError         = errors.New("database operation failed")
)

type registrationService struct {
	repo postgres.UserRepository
}

func NewRegistrationService(repo postgres.UserRepository) RegistrationService {
	return &registrationService{repo: repo}
}

func (s *registrationService) RegisterPersonalInfo(ctx context.Context, user *models.User) (int64, error) {

	if err := validator.ValidateUserRegistration(user); err != nil {
		return 0, errors.Join(ErrInvalidData, err)
	}

	existing, err := s.repo.GetUserByPESEL(ctx, user.Pesel)

	if err != nil {
		return 0, err
	}

	if existing != nil {
		return 0, ErrUserAlreadyRegistered
	}

	id, err := s.repo.CreateUser(ctx, user)

	if err != nil {
		return 0, err
	}

	return id, nil
}


func (s *registrationService) AddResidentialAddress(ctx context.Context, pesel string, addr *models.Address) error {

	user, err := s.repo.GetUserByPESEL(ctx, pesel)

	if err != nil {
		return ErrDatabaseError
	}

	if user == nil {
		return ErrUserNotFound
	}

	if err := validator.ValidateAddress(addr); err != nil {
		return errors.Join(ErrInvalidAddressData, err)
	}

	addr.UserID = user.ID

	if err := s.repo.CreateAddress(ctx, addr); err != nil {
		return ErrDatabaseError
	}

	return nil
}