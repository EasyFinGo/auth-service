package database

import (
	"EasyFinGo/internal/app/auth/model"
	"EasyFinGo/internal/pkg/config"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	LogDataBaseConnected   = "Database connected successfully"
	LogMigrationStandard   = "Running auto migrations"
	LogMigrationsCompleted = "Migrations completed successfully"

	ErrDataBaseConnection = "failed to connect to database"
	ErrDataBaseInstance   = "failed to get database instance"
	ErrMigrationFailed    = "failed to run migration"
)

func NewPostgresDB(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.DBName,
		cfg.Database.SSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	})

	if err != nil {
		return nil, fmt.Errorf("%s: %w", ErrDataBaseConnection, err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", ErrDataBaseInstance, err)
	}

	sqlDB.SetMaxIdleConns(cfg.Database.MaxIdleConns)
	sqlDB.SetMaxOpenConns(cfg.Database.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(cfg.Database.ConnMaxLifeTime)

	log.Println(LogDataBaseConnected)
	return db, nil
}

func AutoMigrate(db *gorm.DB) error {
	log.Println(LogMigrationStandard)

	err := db.AutoMigrate(
		&model.User{},
		&model.Address{},
		&model.Country{},
		&model.Document{},
		&model.Photo{},
	)

	if err != nil {
		return fmt.Errorf("%s: %w", ErrMigrationFailed, err)
	}

	log.Println(LogMigrationsCompleted)
	return nil
}
