

package models

import (
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "os"
    "github.com/joho/godotenv"
    "fmt"
    "log"
)

var DB *gorm.DB //To be used by models 

func init() {
    //Open the database on init
    var err error
    DB, err = CreateDBConnection()
    if err != nil {
	log.Fatal("Could not connect to database: ", err)
    }
}

func getDB() (*gorm.DB) {
    return DB;
}

func CreateDBConnection() (*gorm.DB, error) {
    /* Create a connection for the global variable db */

    if err := godotenv.Load(); err != nil {
	log.Fatal("Error loading .env file: ", err)
    }
    dbHost := os.Getenv("DB_HOST")
    dbName := os.Getenv("DB_NAME")
    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASS")

    credentials := fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUser, dbPassword, dbHost, dbName)

    DB, err := gorm.Open(mysql.Open(credentials), &gorm.Config{})
    return DB, err
}

