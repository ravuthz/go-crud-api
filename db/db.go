package db

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)
 
var db *gorm.DB
var err error
 
type Movie struct {
   ID          string `json:"id" gorm:"primarykey"`
   Name        string `json:"name"`
   Description string `json:"description"`
}
 
func InitPostgresDB() {
   err = godotenv.Load(".env")

   if err != nil {
       log.Fatal("Error loading .env file", err)
   }

   var (
       host = os.Getenv("DB_HOST")
       port = os.Getenv("DB_PORT")
       username = os.Getenv("DB_USERNAME")
       database = os.Getenv("DB_DATABASE")
       password = os.Getenv("DB_PASSWORD")
   )

   dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
       host,
       port,
       username,
       database,
       password,
   )
 
   db, err = gorm.Open("postgres", dsn)

   if err != nil {
       log.Fatal(err)
   }
   
   db.AutoMigrate(Movie{})
}