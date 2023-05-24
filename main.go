package main

import (
	"context"
	"github.com/deliveryhero/pd-order-placement/internal"
	"log"
	"os"
	"path/filepath"
)

func main() {
	payload, err := os.ReadFile(filepath.Join("dataset", "product_payload.json"))
	if err != nil {
		log.Fatal(err)
	}

	validator := &internal.Validator{}
	datastore := &internal.Datastore{}
	ctx := context.Background()

	orderCodeGenerator := internal.NewCodeGenerator()

	payloadStr, err := internal.GetPayloadSkeleton(payload)
	if err != nil {
		log.Fatal(err)
	}

	if vErr := validator.IsDataValid(payloadStr); vErr != nil {
		log.Fatal(vErr)
	}

	code, cErr := orderCodeGenerator.GetOrderCode(payloadStr.Type)
	if cErr != nil {
		log.Fatal(cErr)
	}

	dError := datastore.Flush(ctx, payloadStr, code)
	if dError != nil {
		log.Fatal(dError)
	}
}
