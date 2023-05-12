package main_test

import (
	"github.com/deliveryhero/pd-order-placement/blocks"
	"os"
	"path/filepath"
	"testing"
)

func TestValidPayload(t *testing.T) {

	var tests = map[string]struct {
		PayloadFileName string
	}{
		"payload_with_product": {
			PayloadFileName: "product_payload.json",
		},
	}

	v := &blocks.Validator{}

	for _, test := range tests {
		payload, err := os.ReadFile(filepath.Join("dataset", "valid", test.PayloadFileName))
		if err != nil {
			t.Fatal(err)
		}

		vErr := v.IsDataValid(payload)
		if vErr != nil {
			t.Error("unexpected error: ", err)
		}
	}
}

func TestInvalidPayload(t *testing.T) {

	var tests = map[string]struct {
		PayloadFileName string
		ExpectedError   string
	}{
		"product_item_block_unknown_type": {
			PayloadFileName: "product_item_block_unknown_type.json",
			ExpectedError:   "unknown type",
		},
		"food_product_item_block_invalid": {
			PayloadFileName: "food_product_item_block_invalid.json",
			ExpectedError:   "id field is empty",
		},
		"grocery_product_item_block_invalid": {
			PayloadFileName: "grocery_product_item_block_invalid.json",
			ExpectedError:   "SKU field is empty",
		},
		"payment_block_invalid": {
			PayloadFileName: "payment_block_invalid.json",
			ExpectedError:   "total order value is empty",
		},
		"cross-dependency_food_product_item_related_data_in_payment_section_invalid": {
			PayloadFileName: "cross_dependency_food_product_item_related_data_in_payment_section_invalid.json",
			ExpectedError:   "food item price should be set",
		},
		"cross-dependency_grocery_product_item_related_data_in_payment_section_invalid": {
			PayloadFileName: "cross_dependency_grocery_product_item_related_data_in_payment_section_invalid.json",
			ExpectedError:   "grocery item price should be set",
		},
	}

	v := &blocks.Validator{}

	for _, test := range tests {
		payload, err := os.ReadFile(filepath.Join("dataset", "invalid", test.PayloadFileName))
		if err != nil {
			t.Fatal(err)
		}

		vErr := v.IsDataValid(payload)
		if vErr.Error() != test.ExpectedError {
			t.Error("unexpected error message: ", vErr)
		}
	}
}
