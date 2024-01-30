package main

import (
	"task-intern-product-api/config"
	"task-intern-product-api/routes"
)

func main() {
	config.ConnectDB()
	router := routes.SetupRouter()

	router.Run(":8088")
}