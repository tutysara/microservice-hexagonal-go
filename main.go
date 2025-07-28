package main

import (
	"github.com/tutysara/banking-go/app"
	"github.com/tutysara/banking-go/logger"
)

func main() {
	logger.Info("Starting the application...")
	app.Start()
}
