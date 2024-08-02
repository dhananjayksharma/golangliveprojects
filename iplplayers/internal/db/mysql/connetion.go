package mysql

import (
	"fmt"
	"golangliveprojects/iplplayers/internal/entities"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	MS_USER := os.Getenv("MS_USER_ENV")
	MS_PASS := os.Getenv("MS_PASS_ENV")
	MS_HOST := os.Getenv("MS_HOST_ENV")
	MS_PORT := os.Getenv("MS_PORT_ENV")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/ipldbs?charset=utf8mb4&parseTime=True&loc=Local", MS_USER, MS_PASS, MS_HOST, MS_PORT)
	var err error
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	fmt.Println("DB:", DB)
	// Auto Migrate
	DB.AutoMigrate(&entities.Players{})
	DB.AutoMigrate(&entities.Stadiums{})
	return DB
}
