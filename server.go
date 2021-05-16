package main

import (
	"fmt"
	"graphql_todos/Config"
	"graphql_todos/Router"
	"graphql_todos/graph/model"

	"github.com/jinzhu/gorm"
)

const defaultPort = "8080"

var err error
func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	
	if err != nil {
		fmt.Println("Status: ", err)
	}

	defer Config.DB.Close()
	Config.DB.AutoMigrate(&model.User{})

	r := Router.SetupRouter()
	r.Run()
}
