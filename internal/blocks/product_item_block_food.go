package blocks

import (
	"context"
	"database/sql"
	"errors"
)

type ProductItemFood struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (f ProductItemFood) IsBlockDataValid() error {
	if f.ID == "" {
		return errors.New("id field is empty")
	}

	if f.Name == "" {
		return errors.New("name field is empty")
	}

	return nil
}

func (f ProductItemFood) IsPaymentSectionValid(payment Payment) error {
	if payment.FoodItemPrice == "" {
		return errors.New("food item price should be set")
	}

	return nil
}

func (f ProductItemFood) Persist(ctx context.Context, tx *sql.Tx, orderCode string) error {
	query := "INSERT INTO `product_item` (`order_code`, `id`, `name`) VALUES (?, ?, ?)"
	_, err := tx.ExecContext(ctx, query, orderCode, f.ID, f.Name)

	return err
}
