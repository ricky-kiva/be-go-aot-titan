package main

import (
	"github.com/gin-gonic/gin"
)

// run by `go run .`
func main() {
	r := gin.Default()
	db, err := connectDB()

	setupRouter(r, db, err)

	defer db.Close()

	r.Run()
}
