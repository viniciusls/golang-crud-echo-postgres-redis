package main

import (
	"crud-echo-postgres-redis/api/users/listener"
	"crud-echo-postgres-redis/router"
)

// @title Go Lang CRUD
// @version 1.0
// @description This is a sample Go Lang CRUD application.

// @contact.name Vinícius Silva
// @contact.url https://www.viniciusls.com.br
// @contact.email vinicius.ls@live.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
// @schemes http
func main() {
	go listener.Run()

	e := router.New()

	e.Logger.Fatal(e.Start(":8080"))
}
