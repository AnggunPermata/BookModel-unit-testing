package main

import (
	"project-tdd/anggunpermata/config"
	"project-tdd/anggunpermata/controller"
	"project-tdd/anggunpermata/model"

	"github.com/labstack/echo"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	e := echo.New()
	db, err := gorm.Open(mysql.Open(config.DB_CONNECTION_STRING))
	db.AutoMigrate(&model.Book{})
	if err != nil {
		panic(err)
	}
	bookModel := model.NewGormBookModel(db)
	bookController := controller.CreateGetBookController(bookModel)
	e.GET("/book", bookController)
	e.Start(":8080")
}
