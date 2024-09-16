package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/focusfind/backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func connectDB() *gorm.DB {
	db_password := os.Getenv("DB_PASSWORD")

	dsn := fmt.Sprintf("host=localhost user=postgres password=%s dbname=focusfind port=5432 sslmode=disable", db_password)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connct to database:", err)
	}
	db.AutoMigrate(&models.Spot{}) // automigrate Spot model
	return db
}

func main() {
	// gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	db := connectDB()

	r.GET("/api/spots", func(ctx *gin.Context) {
		var spots []models.Spot
		db.Find(&spots)
		ctx.JSON(http.StatusOK, spots)
	})

	r.POST("/api/spots", func(ctx *gin.Context) {
		var spot models.Spot
		if err := ctx.ShouldBindJSON(&spot); err != nil {
			db.Create(&spot)
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})

	r.Run()
}
