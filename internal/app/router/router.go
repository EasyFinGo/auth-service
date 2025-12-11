package router

import (
	"EasyFinGo/internal/app/health"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, checker *health.Checker) {
	r.GET("/health", healthCheck)
	r.GET("/health/live", livenessCheck(checker))
	r.GET("/health/ready", readinessCheck(checker))

	v1 := r.Group("/api/v1")
	{
		_ = v1
	}
}

func healthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"status":  "ok",
		"message": "User service is running",
		"service": "user-service",
	})
}

func livenessCheck(checker *health.Checker) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := checker.CheckLiveness(); err != nil {
			c.JSON(503, gin.H{
				"status": "unhealthy",
				"error":  err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"status": "healthy",
		})
	}
}

func readinessCheck(checker *health.Checker) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := checker.CheckReadiness(); err != nil {
			c.JSON(503, gin.H{
				"status": "not ready",
				"reason": "database unavailable",
				"error":  err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"status": "ready",
		})
	}
}