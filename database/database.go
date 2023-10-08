package database

import (
	"fmt"
	"uts/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "root@tcp(127.0.0.1:3306)/kebun_binatang?charset=utf8mb4&parseTime=True&loc=Local"
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Database gagal terhubung.")
	}
	fmt.Println("Database berhasil terhubung.")

	DB = conn
	DB.AutoMigrate(models.User{}, models.Admin{})
}
