package main

import (
	"gin-demo/api/routes"
	"gin-demo/config"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.GetDB()
	server := gin.Default()

	routes.InitRoutes(server, db)

	log.Fatal(server.Run(":3000"))

}
