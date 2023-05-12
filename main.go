package main

import (
	"encoding/json"
	"github.com/deliveryhero/pd-order-placement/validation/blocks"
	"log"
	"os"
	"path/filepath"
)

func main() {
	input, err := os.ReadFile(filepath.Join("dataset", "product_payload.json"))
	if err != nil {
		log.Fatal(err)
	}

	var payload blocks.OrderSkeleton
	if err := json.Unmarshal(input, &payload); err != nil {
		log.Fatal(err)
	}

	br := blocks.GetBlockRepository()

	if err := payload.Payment.IsBlockDataValid(); err != nil {
		log.Fatal(err)
	}

	for _, p := range payload.ProductItems {
		sv := br.GetProductStruct(p["type"].(string), p)

		if err := sv.IsBlockDataValid(); err != nil {
			log.Fatal(p["type"].(string), err)
		}

		if err := sv.IsPaymentSectionValid(payload.Payment); err != nil {
			log.Fatal(p["type"].(string), err)
		}
	}

}
