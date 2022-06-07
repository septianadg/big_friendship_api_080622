package main

import (
	// directory root project go yang kita buat
	"./models" // memanggil package models pada directory models
	"./routes"
)

func main() {

	db := models.SetupDB()
	db.AutoMigrate(&models.User{}, &models.Request_friendship{}, &models.Status_friendship{})

	r := routes.SetupRoutes(db)
	r.Run()
}
