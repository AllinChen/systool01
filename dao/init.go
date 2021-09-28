package dao

import (
	"fmt"
	"log"
	"strconv"

	"github.com/AllinChen/systool01/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func GetDB() *gorm.DB {
	return DB
}

func InitDB() {
	port, _ := strconv.Atoi(config.GlobalConf.Mysql.Port)
	conn := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8mb4", config.GlobalConf.Mysql.User,
		config.GlobalConf.Mysql.Pass, config.GlobalConf.Mysql.Host, port, config.GlobalConf.Mysql.DBName)

	var err error
	DB, err = gorm.Open("mysql", conn)
	if err != nil {
		log.Fatal("init DB error : ", err.Error())
	}

	DB.SingularTable(true)
	DB.Callback().Update().Replace("gorm:update_time_stamp", func(scope *gorm.Scope) {})

}
