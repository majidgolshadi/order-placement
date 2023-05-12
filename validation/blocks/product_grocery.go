package blocks

import (
	"errors"
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

func (g Grocery) IsBlockDataValid() error {
	if g.Sku == "" {
		return errors.New("SKU field is empty")
	}

	return nil
}

func (g Grocery) IsPaymentSectionValid(payment Payment) error {
	if payment.GroceryItemPrice == "" {
		return errors.New("grocery item price should be set")
	}

	return nil
}
