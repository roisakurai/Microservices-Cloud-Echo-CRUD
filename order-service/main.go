package main

import (
	"log"
	"order-service/config"
	"order-service/repositories"
	"order-service/routes"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {

	config.ConnectDB()

	if err := repositories.CreateOrderIndexes(); err != nil {
		log.Println("failed to create order indexes:", err)
	}

	e := echo.New()

	routes.InitRoutes(e)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Order service running on port", port)

	e.Logger.Fatal(e.Start(":" + port))

}
