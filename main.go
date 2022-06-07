package main

import (
	// directory root project go yang kita buat
	"Golang_latihan/big_friendship_api/models" // memanggil package models pada directory models
	"Golang_latihan/big_friendship_api/routes"
)

func main() {

	db := models.SetupDB()
	db.AutoMigrate(&models.User{}, &models.Request_friendship{}, &models.Status_friendship{})

	r := routes.SetupRoutes(db)
	r.Run()
}
