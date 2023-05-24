package blocks

import (
	"context"
	"database/sql"
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

func (g ProductItemGrocery) Persist(ctx context.Context, tx *sql.Tx, orderCode string) error {
	query := "INSERT INTO `product_item_grocery` (`order_code`,`sku`) VALUES (?, ?)"
	_, err := tx.ExecContext(ctx, query, orderCode, g.Sku)

	return err
}
