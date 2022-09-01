package dal

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetConnection(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return db, err
}
