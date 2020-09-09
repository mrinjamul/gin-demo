package main

import (
	"gin-demo/api/routes"
	"log"

	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

func main() {
	server := gin.Default()

	routes.InitRoutes(server)

	log.Fatal(server.Run(":3000"))

}
