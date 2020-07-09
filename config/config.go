package config

import "gin-demo/models"

type Database struct {
	Articles []models.Article
}

var db Database

// GetDB retuns database
func GetDB() *Database {
	var articles []models.Article
	db = Database{
		Articles: articles,
	}
	return &db
}
