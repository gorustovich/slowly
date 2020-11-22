package main

import (
	"github.com/gorustovich/slowly/app"
)

func main() {
	server := app.NewServer()
	server.Setup()
	server.Start()
}
