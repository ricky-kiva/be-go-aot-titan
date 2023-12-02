package handler

import (
	"be-go-aot-titan/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func PutHandler(c *gin.Context, db *gorm.DB) {
	var soldier = models.TitanSoldier{}

	soldierId := c.Param("id")

	// could also use `db.Find(&soldier, "id=?", soldierId)` to specify target by `field`
	if db.Find(&soldier, soldierId).RecordNotFound() {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "error",
			"message": "Data not found",
		})
		return
	}

	if c.Bind(&soldier) == nil {
		// could also use `db.Model(&soldier).Where("id=?", soldierId).Update(soldier)` to specify target by `field`
		db.Model(&soldier).Update(soldier)
		c.JSON(http.StatusOK, gin.H{
			"status": "success",
			"data":   soldier,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Bind error",
		})
	}
}
