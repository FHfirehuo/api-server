package model

import (
	"apiserver/pkg/logger"
	"fmt"
	"github.com/spf13/viper"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	Self   *gorm.DB
	Docker *gorm.DB
}

var DB *Database

func open(username, password, addr, name string) *gorm.DB {

	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s",
		username,
		password,
		addr,
		name,
		true,
		"Local")

	//config := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(config), &gorm.Config{})
	if err != nil {
		logger.Exception("failed to connect database", err)
	}

	return db
}

func InitSelfDB() *gorm.DB {
	return open(viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.addr"),
		viper.GetString("db.name"))
}

func GetSelfDB() *gorm.DB {
	return InitSelfDB()
}

func InitDockerDB() *gorm.DB {
	return open(viper.GetString("docker.username"),
		viper.GetString("docker.password"),
		viper.GetString("docker.addr"),
		viper.GetString("docker.name"))
}

func GetDockerDB() *gorm.DB {
	return InitDockerDB()
}

func (db *Database) Init() {

	DB = &Database{
		Self:   GetSelfDB(),
		Docker: GetDockerDB(),
	}
}
