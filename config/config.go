package config

import (
	"fmt"
	"os"
	"time"

	"github.com/astaxie/beego/orm"
)

var db orm.Ormer

// GetDB retuns database
func GetDB() orm.Ormer {
	if db == nil {
		dbHost := os.Getenv("MYSQL_ROOT_HOST")
		dbName := os.Getenv("MYSQL_DATABASE")
		dbUser := os.Getenv("MYSQL_USER")
		dbPassword := os.Getenv("MYSQL_PASSWORD")
		dbPort := os.Getenv("MYSQL_PORT")

		if dbHost == "" {
			fmt.Println("Environment variable DB_HOST is null.")
			return nil
		}
		if dbName == "" {
			fmt.Println("Environment variable DB_NAME is null.")
			return nil
		}
		if dbUser == "" {
			fmt.Println("Environment variable DB_USERNAME is null.")
			return nil
		}
		if dbPassword == "" {
			fmt.Println("Environment variable DB_PASSWORD is null.")
			return nil
		}

		if dbPort == "" {
			dbPort = "3306"
		}
		orm.RegisterDriver("mysql", orm.DRMySQL)
		// DefaultTimeLoc is for timezone it's not mandatory for database connection
		orm.DefaultTimeLoc = time.UTC
		// For mysql need to add network type (tcp/unix) for remote connection must use tcp
		orm.RegisterDataBase("default", "mysql", dbUser+":"+dbPassword+"@tcp("+dbHost+":"+dbPort+")/"+dbName+"?charset=utf8")
		orm.RunSyncdb("default", false, true)
		db = orm.NewOrm()
		db.Using("default")
	}
	return db
}
