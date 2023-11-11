package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=b2b-ms.cry4cich1xjp.ap-southeast-3.rds.amazonaws.com user=postgres password=postgres dbname=postgres port=5432"

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
