package handlers

import "github.com/gin-gonic/gin"

type UserHandler interface {
	RegisterPersonalInfo(c *gin.Context)
	AddResidentialAddress(c *gin.Context)
}