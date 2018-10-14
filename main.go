package main

import "github.com/skanehira/mockapi/app"

func main() {
	app := app.New()
	app.Setup()

	app.Run()
}
