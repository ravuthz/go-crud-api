package main

import (
	"go-crud-api/db"
	"go-crud-api/router"
)
 
func main() {
   db.InitPostgresDB()
   router.InitRouter().Run()
}