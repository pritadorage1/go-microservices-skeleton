package model

import (
	"fmt"

	"go_microservices_skeleton/config"
	"github.com/jinzhu/gorm"
	//we can change this as per your database 
	_ "github.com/jinzhu/gorm/dialects/mysql"
)


func CreateConnection(conf config.DBConfig) (*gorm.DB, error) {
	//we can change this as per your database 
	import (
		_ "github.com/jinzhu/gorm/dialects/mysql"
	)
	// MYSQL Connection string
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		conf.Username, conf.Password, conf.Host, conf.Port, conf.DbName)
	return gorm.Open("mysql", connectionString)

}