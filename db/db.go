package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

var Con *gorm.DB

func Connect() *gorm.DB {
	if Con != nil {
		return Con
	}
	dbURL := os.Getenv(LocoDatabaseUrlkey)
	fmt.Println("Start Time DB Operations")
	fmt.Println(dbURL)
	if dbURL != "" {
		db, err := gorm.Open(mysql.Open(dbURL), &gorm.Config{})
		if err != nil {
			log.Fatalf("Got error when connect database, the error is '%v'", err)
		}
		//db.AutoMigrate(&models.TransactionSummary{})
		Con = db
		return db
	} else {
		fmt.Println("DB Connection here")
		db, err := gorm.Open(mysql.Open("root:12345678@tcp(localhost:3306)/tiger?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
		if err != nil {
			log.Fatalf("Got error jwhen connect database, the error is '%v'", err)
		}
		//db.AutoMigrate(&models.TransactionSummary{})
		Con = db
		return db
	}
}
