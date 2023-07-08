package db

import (
	"log"

	"github.com/athunlal/bookNow-auth-svc/pkg/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func Init(url string) Handler {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&domain.User{})

	return Handler{db}
}
