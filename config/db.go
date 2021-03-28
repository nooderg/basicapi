package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// InitDB connects the DB, returns an error if the connection fails
func InitDB() (*gorm.DB, error) {
	dsn := "root:password@tcp(127.0.0.1:3306)/basic_api?parseTime=True"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
