package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"reflect"
	"strings"
	"tigerhall/models"
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
		//todo: Update DB Creds in the Environment File
		db, err := gorm.Open(mysql.Open(dbURL), &gorm.Config{})
		if err != nil {
			log.Fatalf("Got error when connect database, the error is '%v'", err)
		}
		//db.AutoMigrate(&models.TransactionSummary{})
		Con = db
		return db
	} else {
		fmt.Println("DB Connection here")
		/*
			todo: Update DB Creds in the Environment File
		*/
		db, err := gorm.Open(mysql.Open("root:12345678@tcp(localhost:3306)/tiger?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
		if err != nil {
			log.Fatalf("Got error when connect database, the error is '%v'", err)
		}
		db.AutoMigrate(
			&models.Tiger{},
			&models.AccessToken{},
			&models.User{},
			&models.TigerSight{},
		)
		Con = db
		return db
	}
}

func getModels() []interface{} {
	var models []interface{}
	modelsPkg := reflect.TypeOf(models)
	for i := 0; i < modelsPkg.NumField(); i++ {
		fieldType := modelsPkg.Field(i).Type
		if fieldType.Kind() == reflect.Struct {
			if strings.ToUpper(string(fieldType.Name()[0])) == string(fieldType.Name()[0]) {
				models = append(models, reflect.New(fieldType).Elem().Interface())
			}
		}
	}

	return models
}
