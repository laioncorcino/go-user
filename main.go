package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/laioncorcino/go-user/config/routes"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	fmt.Println(os.Getenv("TEST"))

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup)
	if err := router.Run(":8081"); err != nil {
		log.Fatal(err)
	}
}
