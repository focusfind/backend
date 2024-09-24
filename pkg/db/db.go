package db

import (
	"fmt"
	"log"
	"os"

	"github.com/focusfind/backend/pkg/models"
	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	db_password := os.Getenv("DB_PASSWORD")

	dsn := fmt.Sprintf("postgres://postgres:%s@localhost:5432/focusfind?sslmode=disable", db_password)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("Failed to connct to database:", err)
	}

	db.AutoMigrate(&models.Spot{}, &models.Account{}) // automigrate Spot and Account model

	return db
}
