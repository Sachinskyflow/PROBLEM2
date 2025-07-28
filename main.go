package main

import (
	"fmt"

	app "problem2/pkg/app"
)

func main() {
	if err := app.Main(); err != nil {
		fmt.Printf("App exited with error: %v\n", err)
	}
}
