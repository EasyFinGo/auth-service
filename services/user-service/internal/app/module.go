package app

import (
	"EasyFinGo/internal/config"
	"EasyFinGo/internal/database"
	"EasyFinGo/internal/health"
	"EasyFinGo/internal/router"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"net"
	"net/http"
	"time"
)

var Module = fx.Options(
	fx.Provide(config.LoadConfig),
	fx.Provide(database.NewPostgresDB),
	fx.Provide(health.NewChecker),
	fx.Provide(newRouter),
	fx.Invoke(router.SetupRoutes),
	fx.Invoke(registerHooks),
)

func newRouter(cfg *config.Config) *gin.Engine {
	if cfg.Server.Env == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	return r
}

func registerHooks(
	lc fx.Lifecycle,
	cfg *config.Config,
	router *gin.Engine,
	db *gorm.DB,
	logger *zap.Logger,
) {
	addr := fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port)
	srv := &http.Server{
		Addr:         addr,
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			listener, err := net.Listen("tcp", addr)
			if err != nil {
				return fmt.Errorf("failed to bind to %s: %w", addr, err)
			}
			logger.Info("Starting server",
				zap.String("address", addr),
				zap.String("environment", cfg.Server.Env))

			go func() {
				if err := srv.Serve(listener); err != nil {
					logger.Fatal("failed to start server", zap.Error(err))
				}
			}()
			return nil
		},

		OnStop: func(ctx context.Context) error {
			logger.Info("Shutting down server gracefully")

			shutDownCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancel()

			if err := srv.Shutdown(shutDownCtx); err != nil {
				logger.Error("Server shutdown error", zap.Error(err))
			}
			sqlDB, err := db.DB()
			if err != nil {
				return err
			}

			if err := sqlDB.Close(); err != nil {
				logger.Error("Error closing database", zap.Error(err))
				return err
			}
			logger.Info("database connection closed")
			return nil
		},
	})
}
