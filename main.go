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

	if !payload.Payment.IsBlockDataValid() {
		log.Fatal("Payment section data is not valid")
	}

	for _, p := range payload.ProductItems {
		sv := br.GetProductStruct(p["type"].(string), p)

		if !sv.IsBlockDataValid() {
			log.Fatal(p["type"].(string), " data block is not valid")
		}

		if !sv.IsPaymentSectionValid(payload.Payment) {
			log.Fatal(p["type"].(string), " related payment section data is not valid")
		}
	}

}
