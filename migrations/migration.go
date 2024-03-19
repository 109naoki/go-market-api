package main

import (
	"gin-market/infra"
	"gin-market/models"
)

func main(){
	infra.Initialize()
	db := infra.SetupDB()
	if err := db.AutoMigrate(&models.Item{},&models.User{}); err != nil {
		panic("Failed to migrate database")
	}
}