package main

import (
	"gin-demo/api/routes"
	"log"

	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	server := gin.Default()

	routes.InitRoutes(server)

	log.Fatal(server.Run(":3000"))

}
