package handler

import (
	"be-go-aot-titan/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func GetAllHandler(c *gin.Context, db *gorm.DB) {
	var soldiers []models.TitanSoldier

	db.Find(&soldiers)
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   soldiers,
	})
}

func GetByIdHandler(c *gin.Context, db *gorm.DB) {
	var soldier models.TitanSoldier

	soldierId := c.Param("id")

	if db.Find(&soldier, soldierId).RecordNotFound() {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "error",
			"message": "Data not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    soldier,
	})
}
