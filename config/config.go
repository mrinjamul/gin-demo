package config

import (
	"github.com/astaxie/beego/orm"
)

var db orm.Ormer

// GetDB retuns database
func GetDB() orm.Ormer {
	if db == nil {
		orm.RegisterDriver("mysql", orm.DRMySQL)
		orm.RegisterDataBase("default", "mysql", "root:admin@/gin_demo?charset=utf8")
		orm.RunSyncdb("default", false, true)
		db = orm.NewOrm()
		db.Using("default")
	}
	return db
}
