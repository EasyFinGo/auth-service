package models

import "time"

type User struct {
    ID                  int64     `json:"id"`
    FirstName           string    `json:"first_name" validate:"required,min=1,max=100"`
    MiddleName          *string   `json:"middle_name,omitempty" validate:"omitempty,max=100"`
    LastName            string    `json:"last_name" validate:"required,min=1,max=100"`
    DateOfBirth         time.Time `json:"date_of_birth" validate:"required"`
    Pesel               string    `json:"pesel" validate:"required,len=11,numeric,pesel"` 
    RegistrationStatus  string    `json:"registration_status,omitempty"`
    CreatedAt           time.Time `json:"created_at"`
    UpdatedAt           time.Time `json:"updated_at"`
}

type UserAddress struct {
    ID            int64     `json:"id"`
    UserID        int64     `json:"user_id"`
    Postcode      string    `json:"postcode" validate:"required,postcode_pl"`
    Street        string    `json:"street" validate:"required,min=2"`
    BuildingNo    string    `json:"building_no" validate:"required"`
    ApartmentNo   *string   `json:"apartment_no,omitempty"`
    City          string    `json:"city" validate:"required"`
    Country       string    `json:"country" validate:"required,eq=PL"` 
    CreatedAt     time.Time `json:"created_at"`
}

type UserConsent struct {
	ID          int64     `json:"id"`
	UserID      int64     `json:"user_id"`
	ConsentType string    `json:"consent_type"` //privacy policy, terms of use
	Version     string    `json:"version"`
	AcceptedAt  time.Time `json:"accepted_at"`
	IpAddress   *string   `json:"ip_address,omitempty"`
}
