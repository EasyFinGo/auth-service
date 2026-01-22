package main

import (
	"EasyFinGo/internal/app/auth/models"
	"EasyFinGo/internal/app/auth/repositories/postgres"
	"EasyFinGo/internal/db"
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	db.Init()
	defer db.Close()

	repo := postgres.NewUserRepository(db.DB)

	ctx := context.Background()

	// Example: check if some PESEL already exists
	existing, err := repo.GetUserByPESEL(ctx, "12345678901")
	if err != nil {
		log.Printf("error checking pesel: %v", err)
	}
	if existing != nil {
		fmt.Printf("User already exists: %s %s\n", existing.FirstName, existing.LastName)
		return
	}

	// Example insert (for testing — remove later)
	user := &models.User{
		FirstName:   "Gunel",
		LastName:    "Test",
		DateOfBirth: time.Date(1995, 5, 15, 0, 0, 0, 0, time.UTC),
		Pesel:       "95551512345", // fake PESEL
		IsAccepted:  true,
	}

	id, err := repo.CreateUser(ctx, user)
	if err != nil {
		log.Fatalf("failed to create user: %v", err)
	}

	fmt.Printf("Created user with ID: %d\n", id)

	// You can now also test address creation etc.
}
