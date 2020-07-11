package repository

import (
	"errors"
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

	qs := repo.db.QueryTable(new(models.Article))
	qs = qs.Filter("id", article.ID)

	err := qs.One(&res)
	if err != nil {
		log.Fatal(err)
		return res, err
	}

	return res, nil
}

func (repo *articleRepo) Update(ctx *gin.Context, article *models.Article) error {

	var column []string
	if len(article.Title) > 0 {
		column = append(column, "title")
	}
	if len(article.Description) > 0 {
		column = append(column, "description")
	}
	if len(column) == 0 {
		return errors.New("No changes found")
	}
	_, err := repo.db.Update(article, column...)

	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

func (repo *articleRepo) Delete(ctx *gin.Context, article *models.Article) error {

	qs := repo.db.QueryTable(new(models.Article))
	qs = qs.Filter("id", article.ID)

	_, err := qs.Delete()
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (repo *articleRepo) FindAll(ctx *gin.Context) ([]models.Article, error) {

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
