package main

import (
	"log"

	"github.com/Brun0Santos/api-go-testes/src/controller/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env")
	}

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup)
	if err := router.Run(":9090"); err != nil {
		log.Fatal(err)
	}
}
