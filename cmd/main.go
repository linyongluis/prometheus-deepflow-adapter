package main

import (
	"os"
	"prometheus_adapter/cmd/app"
)

func main() {
	command := app.NewAdapterCommand()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}