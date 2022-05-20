package main

import (
	"crud-echo-postgres-redis/router"
)

func main() {
	e := router.New()

	e.Logger.Fatal(e.Start(":8080"))
}
