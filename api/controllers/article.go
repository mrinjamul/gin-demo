package controllers

import (
	"encoding/json"
	"gin-demo/models"
	"gin-demo/repository"
	"io/ioutil"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Article has operations
type Article interface {
	FindAll(ctx *gin.Context)
	Add(ctx *gin.Context)
	Detail(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
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
	bytes, err := ioutil.ReadAll(ctx.Request.Body)

	if err != nil {
		log.Fatal(err)
	}
	var article models.Article
	err = json.Unmarshal(bytes, &article)
	if err != nil {
		log.Fatal(err)
	}
	svc.articleRepo.Add(ctx, &article)

	ctx.JSON(200, gin.H{
		"message": "Success",
		"article": article,
	})
}

func (svc *article) Detail(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		log.Fatal(err)
	}
	article := models.Article{
		ID: id,
	}

	article, err = svc.articleRepo.Detail(ctx, article)

	ctx.JSON(200, gin.H{
		"message": "Success",
		"article": article,
	})
}

func (svc *article) Update(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		log.Fatal(err)
	}

	bytes, err := ioutil.ReadAll(ctx.Request.Body)

	if err != nil {
		log.Fatal(err)
	}
	var article models.Article
	err = json.Unmarshal(bytes, &article)
	if err != nil {
		log.Fatal(err)
	}
	article.ID = id
	svc.articleRepo.Update(ctx, &article)

	ctx.JSON(200, gin.H{
		"message": "Success",
		"article": article,
	})
}

func (svc *article) Delete(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		log.Fatal(err)
	}
	article := models.Article{
		ID: id,
	}

	svc.articleRepo.Delete(ctx, &article)

	ctx.JSON(200, gin.H{
		"message": "Success",
		"article": article,
	})
}

// NewArticle new article
func NewArticle(
	articleRepo repository.ArticleRepo,
) Article {
	return &article{
		articleRepo: articleRepo,
	}
}
