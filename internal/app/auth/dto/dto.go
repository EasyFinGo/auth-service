package dto

import "time"

type RegisterPersonalInfoRequest struct {
	FirstName   string    `json:"first_name" binding:"required,min=1"`
	LastName    string    `json:"last_name" binding:"required,min=1"`
	MiddleName  *string   `json:"middle_name,omitempty"`
	DateOfBirth time.Time `json:"date_of_birth" binding:"required"`
	PESEL       string    `json:"pesel" binding:"required,len=11,numeric"`
	IsAccepted  bool      `json:"is_accepted" binding:"required"`
}

type RegisterPersonalInfoResponse struct {
	ID        int64     `json:"id"`
	PESEL     string    `json:"pesel"`
	CreatedAt time.Time `json:"created_at"`
	Message   string    `json:"message,omitempty"`
}

type AddAddressRequest struct {
	Postcode    string  `json:"postcode" binding:"required"`
	Street      string  `json:"street" binding:"required,min=2"`
	BuildingNo  string  `json:"building_no" binding:"required"`
	ApartmentNo *string `json:"apartment_no,omitempty"`
	City        string  `json:"city" binding:"required,min=2"`
	Country     string  `json:"country,omitempty"`
}

type SuccessResponse struct {
	Message string `json:"message"`
}
