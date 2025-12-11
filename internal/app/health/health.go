package health

import (
	"context"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

type Checker struct {
	db     *gorm.DB
	logger *zap.Logger
}

func NewChecker(db *gorm.DB, logger *zap.Logger) *Checker {
	return &Checker{
		db:     db,
		logger: logger,
	}
}

func (c *Checker) CheckLiveness() error {
	return nil
}

func (c *Checker) CheckReadiness() error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	sqlDb, err := c.db.DB()
	if err != nil {
		c.logger.Error("failed to get database", zap.Error(err))
		return err
	}

	if err := sqlDb.PingContext(ctx); err != nil {
		c.logger.Error("database ping failed")
	}
	return nil
}