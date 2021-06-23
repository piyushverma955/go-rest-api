package main

import (
	"fmt"
	"go-rest-api/Config"
	"go-rest-api/Models"
	"go-rest-api/Routes"

	"github.com/jinzhu/gorm"
)

var err error

func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildConfig()))

	if err != nil {
		fmt.Println(err)
	}

	defer Config.DB.Close()

	Config.DB.AutoMigrate(&Models.User{})

	r := Routes.SetUpRouter()

	r.Run()
}
