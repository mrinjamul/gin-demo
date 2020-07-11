package repository

import (
	"gin-demo/models"
	"log"

	"github.com/astaxie/beego/orm"

	"github.com/gin-gonic/gin"
)

type ArticleRepo interface {
	Add(ctx *gin.Context, article *models.Article) error
	FindAll(ctx *gin.Context) ([]models.Article, error)
	Detail(ctx *gin.Context, article models.Article) (models.Article, error)
	Update(ctx *gin.Context, article *models.Article) error
	Delete(ctx *gin.Context, article *models.Article) error
}

type articleRepo struct {
	db orm.Ormer
}

func (repo *articleRepo) Add(ctx *gin.Context, article *models.Article) error {
	id, err := repo.db.Insert(article)
	if err != nil {
		return err
	}
	article.ID = id
	return nil
}

func (repo *articleRepo) Detail(ctx *gin.Context, article models.Article) (models.Article, error) {

	var res models.Article

	// for _, a := range repo.db.Articles {
	// 	if a.ID == article.ID {
	// 		res = a
	// 		break
	// 	}
	// }

	return res, nil
}

func (repo *articleRepo) Update(ctx *gin.Context, article *models.Article) error {

	// articles := repo.db.Articles

	// var index int

	// var oldArticle models.Article

	// for i, a := range articles {
	// 	if a.ID == article.ID {
	// 		oldArticle = a
	// 		index = i
	// 		break
	// 	}
	// }

	// if len(article.Title) > 0 {
	// 	oldArticle.Title = article.Title
	// }
	// if len(article.Description) > 0 {
	// 	oldArticle.Description = article.Description
	// }
	// articles[index] = oldArticle
	// repo.db.Articles = articles
	// article = &oldArticle
	return nil
}

func (repo *articleRepo) Delete(ctx *gin.Context, article *models.Article) error {

	// articles := repo.db.Articles
	// var index int
	// var oldArticle models.Article

	// for i, a := range articles {
	// 	if a.ID == article.ID {
	// 		oldArticle = a
	// 		index = i
	// 		break
	// 	}
	// }

	// repo.db.Articles = append(articles[:index], articles[index+1:]...)
	// article = &oldArticle
	return nil
}

func (repo *articleRepo) FindAll(ctx *gin.Context) ([]models.Article, error) {

	// articles := repo.db.Articles

	// return articles, nil

	qs := repo.db.QueryTable(new(models.Article))
	var articles []models.Article
	_, err := qs.All(&articles)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return articles, nil
}

func NewArticleRepo(db orm.Ormer) ArticleRepo {
	return &articleRepo{
		db: db,
	}
}
