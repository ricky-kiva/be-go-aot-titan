package handler

import (
	"be-go-aot-titan/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func DelHandler(c *gin.Context, db *gorm.DB) {
	var soldier models.TitanSoldier

	soldierId := c.Param("id")
	db.Delete(&soldier, soldierId)

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Delete success",
	})
}