package main

import (
	"github.com/deliveryhero/pd-order-placement/blocks"
	"log"
	"os"
	"path/filepath"
)

func main() {
	payload, err := os.ReadFile(filepath.Join("dataset", "product_payload.json"))
	if err != nil {
		log.Fatal(err)
	}

	validator := &blocks.Validator{}
	vErr := validator.IsDataValid(payload)

	if vErr != nil {
		log.Fatal(vErr)
	}

}
