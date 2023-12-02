package handler

import (
	"be-go-aot-titan/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func PostHandler(c *gin.Context, db *gorm.DB) {
	var newSoldier models.TitanSoldier

	if c.Bind(&newSoldier) == nil {
		db.Create(&newSoldier)
		c.JSON(http.StatusOK, gin.H{
			"status": "success",
			"data":   newSoldier,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Bind error",
		})
	}
}
