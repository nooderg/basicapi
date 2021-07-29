package config

import (
	"basic-api/models"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


var DBClient *gorm.DB

// InitDB connects the DB, returns an error if the connection fails
func InitDB() {
	for {
		log.Println("Trying to connect to the DB...")
		dsn := "root:password@tcp(127.0.0.1:3306)/basic_api?parseTime=True"
		client, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Println("Error connecting to DB, retrying in 5 seconds...")
			time.Sleep(time.Second * 5)
		}

		fillDB(client)

		DBClient = client
		return
	}
}

func fillDB(client *gorm.DB) {
	for {
		err := client.AutoMigrate(&models.User{}, &models.Article{}, &models.Comment{}, &models.Opinion{})
		if err != nil {
			log.Println("Error connecting to DB, retrying in 5 seconds...")
			continue
		}
		return
	}
}