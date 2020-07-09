package repository

import (
	"gin-demo/config"
	"gin-demo/models"

	"github.com/gin-gonic/gin"
)

type ArticleRepo interface {
	Add(ctx *gin.Context, article *models.Article) error
	FindAll(ctx *gin.Context) ([]models.Article, error)
}

type articleRepo struct {
	db *config.Database
}

func (repo *articleRepo) Add(ctx *gin.Context, article *models.Article) error {
	articles := repo.db.Articles
	if articles == nil {
		articles = []models.Article{}
	}
	articles = append(articles, *article)
	repo.db.Articles = articles
	return nil
}

func (repo *articleRepo) FindAll(ctx *gin.Context) ([]models.Article, error) {

	articles := repo.db.Articles

	return articles, nil
}

func NewArticleRepo(db *config.Database) ArticleRepo {
	return &articleRepo{
		db: db,
	}
}
