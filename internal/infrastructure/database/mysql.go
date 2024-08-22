package infrastructure

import (
	"fiber-gorm/internal/config"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseInit() {
	config.InitConfig()

	dsn := config.GetEnv("DB_USERNAME") + ":" + config.GetEnv("DB_PASSWORD") + "@tcp(" + config.GetEnv("DB_HOST") + ":" + config.GetEnv("DB_PORT") + ")/" + config.GetEnv("DB_DATABASE") + "?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}
}
