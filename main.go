package main

import (
	"fmt"
	"log"
	"github.com/Prototype-1/loyalty-points-system/config"
	"github.com/Prototype-1/loyalty-points-system/database"
	"github.com/Prototype-1/loyalty-points-system/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	c := config.LoadConfig()
	database.ConnectDatabase(c)

	r := gin.Default()
	routes.SetupRoutes(r)

	port := c.ServerPort
	fmt.Printf("Server is running on port %s...\n", port)
	log.Fatal(r.Run(":" + port))
}
