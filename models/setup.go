package models

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func ConnectDatabase() {
	DBHost := "localhost"
	DBPort := "5432"
	DBUser := "postgres"
	DBName := "cardb"
	DBPassword := "postgresql"
	DBDriver := "postgres"
	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DBHost, DBPort, DBUser, "postgres", DBPassword)

	var err error
	DB, _ = gorm.Open(DBDriver, DBURL)
	DB.Exec(fmt.Sprintf("DROP DATABASE IF EXISTS %s;", DBName))
	DB.Exec(fmt.Sprintf("CREATE DATABASE %s;", DBName))
	DB.Close()
	DBURL = fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DBHost, DBPort, DBUser, DBName, DBPassword)
	DB, err = gorm.Open(DBDriver, DBURL)
	if err != nil {
		fmt.Printf("Cannot connect to %s database \n", DBDriver)
		log.Fatal(err)
	}

	DB.AutoMigrate(&Car{})
}
