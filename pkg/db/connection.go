package db

import (
	"fmt"

	"github.com/athunlal/bookNow-auth-svc/pkg/config"
	"github.com/athunlal/bookNow-auth-svc/pkg/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDb(cfg config.Config) (*gorm.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", cfg.DBHost, cfg.DBUser, cfg.DBName, cfg.DBPort, cfg.DBPassword)
	DB, dbErr := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	DB.AutoMigrate(&domain.User{})

	return DB, dbErr

}
