package blocks

import (
	"github.com/mitchellh/mapstructure"
)

type Grocery struct {
	Code string `json:"code,omitempty"`
	Sku  string `json:"sku"`
}

func (g Grocery) MapUnmarshal(data map[string]any) (SectionValidator, error) {
	var grocery Grocery
	err := mapstructure.Decode(data, &grocery)

	return grocery, err
}

func (g Grocery) IsBlockDataValid() bool {
	if g.Sku == "" {
		return false
	}

	return true
}

func (g Grocery) IsPaymentSectionValid(payment Payment) bool {
	if payment.GroceryItemPrice == "" {
		return false
	}

	return true
}
