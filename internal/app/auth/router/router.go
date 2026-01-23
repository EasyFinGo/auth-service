package router

import (
	handlers "EasyFinGo/internal/app/auth/handlers/http"
	"EasyFinGo/internal/app/auth/services"

	"github.com/gin-gonic/gin"
)

func Setup(svc services.RegistrationService) *gin.Engine {
	r := gin.Default()

	api := r.Group("/api/v1")

	userH := handlers.NewUserHandler(svc)

	api.POST("/register/personal", userH.RegisterPersonalInfo)
	api.POST("/register/address/:pesel", userH.AddResidentialAddress)

	return r
}