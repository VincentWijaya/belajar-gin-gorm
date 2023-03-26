package controllers

import (
	"belajar-gin-gorm/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *Controllers) CreateUser(ctx *gin.Context) {
	var (
		request models.CreateUserRequest
	)

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	user := models.User{
		Email: request.Email,
	}

	if err := c.masterDB.Create(&user).Error; err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, user)
	return
}

func (c *Controllers) GetUsersWithProducts(ctx *gin.Context) {
	user := []models.User{}

	err := c.masterDB.Preload("Products").Find(&user).Error
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, user)
	return
}
