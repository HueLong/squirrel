package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
	"strconv"
	"time"
)

var Db *gorm.DB

func Init() {
	dsn := "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dsn = fmt.Sprintf(dsn, user, password, host, port, dbName)
	var config = &gorm.Config{
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
	}
	if os.Getenv("GIN_ENV") != "prod" {
		config.Logger = logger.Default.LogMode(logger.Info)
	}
	db, err := gorm.Open(mysql.Open(dsn), config)
	if err != nil {
		fmt.Println("数据库链接出错:" + err.Error())
		//数据出错直接退出
		os.Exit(1)
	} else {
		maxIdle, _ := strconv.Atoi(os.Getenv("DB_MAX_IDLE"))
		maxOpen, _ := strconv.Atoi(os.Getenv("DB_MAX_OPEN"))
		sqlDB, _ := db.DB()
		sqlDB.SetConnMaxLifetime(time.Minute * 10)
		sqlDB.SetConnMaxIdleTime(time.Minute * 5)
		sqlDB.SetMaxIdleConns(maxIdle)
		sqlDB.SetMaxOpenConns(maxOpen)
		Db = db
	}
}
