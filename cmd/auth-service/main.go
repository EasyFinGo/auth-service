package main

import (
	"EasyFinGo/internal/app/auth/repositories/postgres"
	"EasyFinGo/internal/app/auth/router"
	"EasyFinGo/internal/app/auth/services"
	"EasyFinGo/internal/db"
)

func main() {
	db.Init()
	defer db.Close()

	repo := postgres.NewUserRepository(db.DB)

	svc := services.NewRegistrationService(repo)
	r := router.Setup(svc)
	r.Run(":8080")
}
