package models

import "time"

type User struct {
	ID          int64     `json:"id"`
	FirstName   string    `json:"first_name"`
	MiddleName  *string   `json:"middle_name,omitempty"`
	LastName    string    `json:"last_name"`
	DateOfBirth time.Time `json:"date_of_birth"`
	Pesel       string    `json:"pesel"`
	IsAccepted  bool      `json:"is_accepted"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Address struct {
	ID          int64     `json:"id"`
	UserID      int64     `json:"user_id"`
	Postcode    string    `json:"postcode"`
	Street      string    `json:"street"`
	BuildingNo  string    `json:"building_no"`
	ApartmentNo *string   `json:"apartment_no,omitempty"`
	City        string    `json:"city"`
	Country     string    `json:"country"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
