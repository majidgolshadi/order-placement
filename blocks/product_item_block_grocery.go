package blocks

import (
	"errors"
)

type ProductItemGrocery struct {
	Code string `json:"code,omitempty"`
	Sku  string `json:"sku"`
}

func (g ProductItemGrocery) IsBlockDataValid() error {
	if g.Sku == "" {
		return errors.New("SKU field is empty")
	}

	return nil
}

func (g ProductItemGrocery) IsPaymentSectionValid(payment Payment) error {
	if payment.GroceryItemPrice == "" {
		return errors.New("grocery item price should be set")
	}

	return nil
}
