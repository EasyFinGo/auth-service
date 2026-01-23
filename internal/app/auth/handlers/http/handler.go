package handlers

import (
	"EasyFinGo/internal/app/auth/dto"
	"EasyFinGo/internal/app/auth/models"
	"EasyFinGo/internal/app/auth/services"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	svc services.RegistrationService
}

func NewUserHandler(svc services.RegistrationService) UserHandler {
	return &userHandler{svc: svc}
}

func (h *userHandler) RegisterPersonalInfo(c *gin.Context) {
	var req dto.RegisterPersonalInfoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := &models.User{
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		MiddleName:  req.MiddleName,
		DateOfBirth: req.DateOfBirth,
		Pesel:       req.PESEL,
		IsAccepted:  req.IsAccepted,
	}

	id, err := h.svc.RegisterPersonalInfo(c.Request.Context(), user)

	if err != nil {
		var status int
		var message string

		switch {
		case errors.Is(err, services.ErrUserAlreadyRegistered):
			status = http.StatusConflict
			message = "User with this PESEL already exists"
		case errors.Is(err, services.ErrInvalidData):
			status = http.StatusBadRequest
			message = "Invalid input data"
		default:
			status = http.StatusInternalServerError
			message = "Internal server error"
		}

		c.JSON(status, gin.H{"error": message})
		return
	}
	resp := dto.RegisterPersonalInfoResponse{
		ID:        id,
		PESEL:     req.PESEL,
		CreatedAt: time.Now(),
		Message:   "Personal information registered successfully",
	}

	c.JSON(http.StatusCreated, resp)
}

func (h *userHandler) AddResidentialAddress(c *gin.Context) {
	pesel := c.Param("pesel")
	if pesel == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "PESEL is required in path"})
		return
	}

	var req dto.AddAddressRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	addr := &models.Address{
		Postcode:    req.Postcode,
		Street:      req.Street,
		BuildingNo:  req.BuildingNo,
		ApartmentNo: req.ApartmentNo,
		City:        req.City,
		Country:     req.Country,
	}

	err := h.svc.AddResidentialAddress(c.Request.Context(), pesel, addr)
	if err != nil {
		var status int
		var message string

		switch {
		case errors.Is(err, services.ErrUserNotFound):
			status = http.StatusNotFound
			message = "User not found"
		case errors.Is(err, services.ErrInvalidAddressData):
			status = http.StatusBadRequest
			message = "Invalid address data"
		default:
			status = http.StatusInternalServerError
			message = "Failed to save address"
		}

		c.JSON(status, gin.H{"error": message})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse{
		Message: "Address successfully added",
	})
}
