package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type dbconfig struct {
	Host     string
	User     string
	Password string
	Port     int
	Dbname   string
}

func Buildconfig() *dbconfig {

	dbconf := dbconfig{
		Host:     "localhost",
		User:     "admin",
		Password: "user",
		Port:     8080,
		Dbname:   "userdata",
	}
	return &dbconf
}

func Buildurl(db *dbconfig) string {

	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		db.User,
		db.Password,
		db.Host,
		db.Port,
		db.Dbname)
}
