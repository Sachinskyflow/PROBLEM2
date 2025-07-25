package main

import (
	"context"
	"fmt"
	"os"

	app "problem2/pkg/app"

	aerospike "github.com/aerospike/aerospike-client-go/v8"
)

const (
	host = "aerospike"
	port = 3000
)

func main() {
	ctx := context.Background()
	aero, err := aerospike.NewClient(host, port)
	if err != nil {
		fmt.Printf("Failed to create aerospike client: %v\n", err)
		os.Exit(1)
	}
	defer aero.Close()
	if err := app.Main(ctx, aero); err != nil {
		fmt.Printf("App exited with error: %v\n", err)
	}
}
