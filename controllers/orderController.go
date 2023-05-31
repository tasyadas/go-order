package controllers

import (
	"assignment-2/database"
	"assignment-2/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOrder(ctx *gin.Context) {
	db := database.GetDB()
	var order models.Order

	if err := ctx.ShouldBindJSON(&order); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}

	errOrder := db.Debug().Create(&order).Error

	if errOrder != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": errOrder.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"order": order,
		"code":  200,
	})
}

func GetData(ctx *gin.Context) {
	db := database.GetDB()
	var orders []models.Order

	err := db.Model(&models.Order{}).Preload("Items").Find(&orders).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"orders": orders,
		"code":   200,
	})
}

func UpdateData(ctx *gin.Context) {
	db := database.GetDB()
	var order models.Order
	id := ctx.Query("id")

	if err := db.First(&models.Order{}, "order_id = ?", id).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"result": fmt.Sprintf("gagal Update order Id %v tidak di temukan", id),
		})
		return
	}

	if err := ctx.ShouldBindJSON(&order); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}

	err := db.Model(&models.Order{}).Where("order_id = ?", id).Updates(order).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"orders": order,
		"code":   200,
	})
}

func DeleteData(ctx *gin.Context) {
	db := database.GetDB()
	order := models.Order{}
	id := ctx.Query("id")

	if err := db.First(&order, "order_id = ?", id).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"result": fmt.Sprintf("gagal Update order Id %v tidak di temukan", id),
		})
		return
	}

	if err := db.Delete(&order).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("order dengan id %v berhasil dihapus", id),
		"code":    200,
	})
}
