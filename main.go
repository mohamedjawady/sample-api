package main

import (
	"context"
	"log"

	"github.com/mohamedjawady/sample-api/application"
)

func main() {
	app := application.NewApp()

	err := app.Start(context.TODO())
	if err != nil {
		log.Println(err)
	}
}
