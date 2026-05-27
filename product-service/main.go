package main

import (
	"log"
	"os"
	"product-service/config"
	"product-service/routes"

	"github.com/labstack/echo/v4"
)

func main() {

	config.ConnectDB()

	e := echo.New()

	routes.InitRoutes(e)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("User service running on port", port)
	e.Logger.Fatal(e.Start(":" + port))

}
