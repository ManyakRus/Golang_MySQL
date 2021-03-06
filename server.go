package main

import (
	"MySQL/infrastructure"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
	//dsn = "Server=127.0.0.1;Database=Test;Uid=root;Pwd=111"
	dsn = "root:111@tcp(127.0.0.1:3306)/Test?charset=utf8mb4&parseTime=True&loc=Local"
)

func main() {
	dbinit()
	infrastructure.Init()
	e := echo.New()
	e.Logger.Fatal(e.Start(":1323"))
}

func dbinit() {
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
	}
	//db.Migrator().CreateTable(domain.User{})
}
