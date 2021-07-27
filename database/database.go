package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func CreateDB(connectionString string) (*gorm.DB, error) {
	DB, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		return DB, err
	}
	return DB, err
}
