package controllers

import (
	"gin-demo/models"
	"gin-demo/repository"

	"github.com/gin-gonic/gin"
)

type Article interface {
	FindAll(ctx *gin.Context)
	Add(ctx *gin.Context)
}

type article struct {
	articleRepo repository.ArticleRepo
}

func (svc *article) FindAll(ctx *gin.Context) {
	articles, _ := svc.articleRepo.FindAll(ctx)

	ctx.JSON(200, gin.H{
		"message":  "Success",
		"articles": articles,
	})
}

func (svc *article) Add(ctx *gin.Context) {
	article := models.Article{
		ID:          1,
		Title:       "My title",
		Description: "My Description",
	}
	svc.articleRepo.Add(ctx, &article)

	ctx.JSON(200, gin.H{
		"message": "Success",
		"article": article,
	})
}

func NewArticle(
	articleRepo repository.ArticleRepo,
) Article {
	return &article{
		articleRepo: articleRepo,
	}
}
