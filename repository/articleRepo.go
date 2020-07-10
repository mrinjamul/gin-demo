package repository

import (
	"gin-demo/config"
	"gin-demo/models"

	"github.com/gin-gonic/gin"
)

type ArticleRepo interface {
	Add(ctx *gin.Context, article *models.Article) error
	FindAll(ctx *gin.Context) ([]models.Article, error)
	Detail(ctx *gin.Context, article models.Article) (models.Article, error)
}

type articleRepo struct {
	db *config.Database
}

func (repo *articleRepo) Add(ctx *gin.Context, article *models.Article) error {
	articles := repo.db.Articles
	if articles == nil {
		articles = []models.Article{}
	}
	var id int64 = 1
	if len(articles) > 0 {
		last := articles[len(articles)-1]
		id = last.ID + 1
	}
	article.ID = id
	articles = append(articles, *article)
	repo.db.Articles = articles
	return nil
}

func (repo *articleRepo) Detail(ctx *gin.Context, article models.Article) (models.Article, error) {

	var res models.Article

	for _, a := range repo.db.Articles {
		if a.ID == article.ID {
			res = a
			break
		}
	}

	return res, nil
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
