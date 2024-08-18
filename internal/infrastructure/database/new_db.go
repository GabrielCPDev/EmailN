package database

import (
	"emailn/internal/domain/campaign"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDb() *gorm.DB {
	dsn := "host=localhost user=user password=123 dbname=emailn-dev port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("fail to connet to database")
	}

	db.AutoMigrate(&campaign.Campaign{}, &campaign.Contact{})
	return db
}
