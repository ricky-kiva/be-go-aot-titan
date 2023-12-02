package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

// run by `go run .`
func main() {
	r := gin.Default()
	db, err := connectDB()

	setupRouter(r, db, err)

	defer db.Close()

	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	r.Run("0.0.0.0:" + port)
}
