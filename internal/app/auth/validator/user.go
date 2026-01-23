package validator

import (
	"EasyFinGo/internal/app/auth/models"
	"errors"
	"regexp"
	"strings"
	"time"
)

var (
	ErrInvalidFirstName   = errors.New("first name must contain only latin letters, spaces, apostrophes or dashes")
	ErrInvalidLastName    = errors.New("last name must contain only latin letters, spaces, apostrophes or dashes")
	ErrUnderage           = errors.New("user must be at least 18 years old")
	ErrInvalidPESEL       = errors.New("PESEL must be exactly 11 digits")
	ErrPrivacyNotAccepted = errors.New("must accept privacy policy and terms")
	ErrInvalidDateOfBirth = errors.New("date of birth is invalid or in the future")
)

var latinNameRegex = regexp.MustCompile(`^[A-Za-z' -]+$`)

func ValidateUserRegistration(u *models.User) error {
	if u == nil {
		return errors.New("user is required")
	}

	first := strings.TrimSpace(u.FirstName)
	if first == "" || len(first) < 1 || !latinNameRegex.MatchString(first) {
		return ErrInvalidFirstName
	}

	last := strings.TrimSpace(u.LastName)
	if last == "" || len(last) < 1 || !latinNameRegex.MatchString(last) {
		return ErrInvalidLastName
	}

	if !u.IsAccepted {
		return ErrPrivacyNotAccepted
	}

	if len(u.Pesel) != 11 || !regexp.MustCompile(`^[0-9]{11}$`).MatchString(u.Pesel) {
		return ErrInvalidPESEL
	}

	if u.DateOfBirth.IsZero() || u.DateOfBirth.After(time.Now()) {
		return ErrInvalidDateOfBirth
	}
	age := int(time.Since(u.DateOfBirth).Hours() / 24 / 365.25)
	if age < 18 {
		return ErrUnderage
	}
	return nil
}

func ValidateAddress(a *models.Address) error {
	if a == nil {
		return errors.New("address is required")
	}

	postcode := strings.TrimSpace(a.Postcode)
	if postcode == "" {
		return errors.New("postcode is required")
	}

	street := strings.TrimSpace(a.Street)
	if street == "" || len(street) < 2 {
		return errors.New("street name is too short or missing")
	}

	if strings.TrimSpace(a.BuildingNo) == "" {
		return errors.New("building number is required")
	}

	city := strings.TrimSpace(a.City)
	if city == "" || len(city) < 2 {
		return errors.New("city is required")
	}

	return nil
}