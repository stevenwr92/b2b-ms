package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=dev-db user=postgres password=123 dbname=go port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&User{})
	DB = db
}

func CloseDatabase() {
	if DB != nil {
		sqlDB, err := DB.DB()
		if err != nil {
			panic(err)
		}
		sqlDB.Close()
	}
}
