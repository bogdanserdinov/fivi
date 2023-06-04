package main

import (
	"context"
	"log"

	gateway_app "fivi/services/gateway/v1/cmd/gateway/app"
)

func main() {
	ctx := context.Background()

	app, err := gateway_app.New()
	if err != nil {
		log.Fatal(err)
	}

	app.Run(ctx)
}
