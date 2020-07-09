package routes

import (
	"gin-demo/api/services"
	"gin-demo/config"

	"github.com/gin-gonic/gin"
)

func InitRoutes(server *gin.Engine, db *config.Database) {
	svc := services.NewServices(db)
	server.GET("/", func(c *gin.Context) {
		svc.ArticleService().FindAll(c)
	})

	server.GET("/add", func(c *gin.Context) {
		svc.ArticleService().Add(c)
	})
}
