package main

import (
	"be-go-aot-titan/auth"
	"be-go-aot-titan/handler"
	"be-go-aot-titan/middleware"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func setupRouter(r *gin.Engine, db *gorm.DB, err error) {
	if err != nil {
		panic("DB connection failed")
	}
	fmt.Println("DB connection success")

	migrateDB(db)

	r.POST("/login", auth.LoginHandler) // can also pass like this, if only need `*gin.Context`
	r.POST("/aot", func(ctx *gin.Context) { handler.PostHandler(ctx, db) })

	// require to login to access these
	r.GET("/aot", middleware.AuthValidation, func(ctx *gin.Context) { handler.GetAllHandler(ctx, db) })
	r.GET("/aot/:id", middleware.AuthValidation, func(ctx *gin.Context) { handler.GetByIdHandler(ctx, db) })

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Dedicate your heart!"})
	})

	r.PUT("/aot/:id", func(ctx *gin.Context) { handler.PutHandler(ctx, db) })

	r.DELETE("/aot/:id", func(ctx *gin.Context) { handler.DelHandler(ctx, db) })
}
