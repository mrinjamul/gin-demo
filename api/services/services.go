package services

import (
	"gin-demo/api/controllers"
	"gin-demo/config"
	"gin-demo/repository"
)

// Services interface of service
type Services interface {
	ArticleService() controllers.Article
}

type services struct {
	article controllers.Article
}

func (svc *services) ArticleService() controllers.Article {
	return svc.article
}

// NewServices initializes service
func NewServices() Services {
	db := config.GetDB()
	return &services{
		article: controllers.NewArticle(
			repository.NewArticleRepo(db),
		),
	}
}
