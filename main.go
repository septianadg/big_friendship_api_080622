package main

import (
	// directory root project go yang kita buat
	"https://github.com/septianppm/big_friendship_api_080622/tree/main/models" // memanggil package models pada directory models
	"https://github.com/septianppm/big_friendship_api_080622/tree/main/routes"
)

func main() {

	db := models.SetupDB()
	db.AutoMigrate(&models.User{}, &models.Request_friendship{}, &models.Status_friendship{})

	r := routes.SetupRoutes(db)
	r.Run()
}
