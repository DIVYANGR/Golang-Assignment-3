package main

import (
	"github.com/Divyang/Todo/app"
	"github.com/Divyang/Todo/config"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":3000")
}
