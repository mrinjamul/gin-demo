package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

// Article is a model which will use to store data in datanase
type Article struct {
	ID          int64     `json:"id" orm:"column(id);pk;auto"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at" orm:"auto_now_add;type(datetime)"`
	UpdatedAt   time.Time `json:"updated_at" orm:"auto_now;type(datetime)"`
}

// TableName : It will returns table name
func (a *Article) TableName() string {
	return "articles"
}

func init() {
	orm.RegisterModel(new(Article))
}
