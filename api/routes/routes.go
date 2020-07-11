package routes

import (
	"gin-demo/api/services"

	"github.com/gin-gonic/gin"
)

func InitRoutes(server *gin.Engine) {
	svc := services.NewServices()
	server.GET("/articles", func(c *gin.Context) {
		svc.ArticleService().FindAll(c)
	})

	server.POST("/articles", func(c *gin.Context) {
		svc.ArticleService().Add(c)
	})

	server.GET("/articles/:id", func(c *gin.Context) {
		svc.ArticleService().Detail(c)
	})

	server.PATCH("/articles/:id", func(c *gin.Context) {
		svc.ArticleService().Update(c)
	})

	server.DELETE("/articles/:id", func(c *gin.Context) {
		svc.ArticleService().Delete(c)
	})
}
