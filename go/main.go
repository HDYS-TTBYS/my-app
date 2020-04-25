package main

import (
	"fmt"
	"main/controller"
	"main/model"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	//db
	db, err := gorm.Open("postgres", "host=postgres port=5432 user=root dbname=go_app_db password=password sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.LogMode(true)
	if err := db.DB().Ping(); err != nil {
		panic(err)
	}
	db.LogMode(true)
	fmt.Println("Db connect!")
	db.AutoMigrate(&model.User{})

	//echo
	e := echo.New()
	defer e.Close()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	http.Handle("/", e)
	g := e.Group("/api/v1")

	controller.NewUser(db).Handle(g)

	e.Start(":8080")

}
