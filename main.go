package main

import (
	"cellify_backend/app"
	"cellify_backend/database"
	"log"
)

func main() {
	db := database.MyDataBase{}
	_, err := db.DataBaseINIT()
	if err != nil {
		log.Fatal(err)
	}
	app.AppRoutes()
}
