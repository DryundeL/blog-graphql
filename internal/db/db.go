package db

import (
	"github.com/DryundeL/blog-graphql/internal/config"
	"github.com/DryundeL/blog-graphql/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func InitDB(cfg *config.Config) *gorm.DB {
	db, err := gorm.Open(postgres.Open(cfg.DBConn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to db:", err)
	}
	err = db.AutoMigrate(&models.User{}, &models.Post{})
	if err != nil {
		return nil
	}

	return db
}
