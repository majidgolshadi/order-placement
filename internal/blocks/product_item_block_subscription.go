package blocks

import (
	"context"
	"database/sql"
	"errors"
)

type ProductItemSubscription struct {
	SubscriptionID string `json:"subscription_id"`
}

func (s ProductItemSubscription) IsBlockDataValid() error {
	if s.SubscriptionID == "" {
		return errors.New("subscription id field is empty")
	}

	return nil
}

func (s ProductItemSubscription) IsPaymentSectionValid(payment Payment) error {
	return nil
}

func (s ProductItemSubscription) Persist(ctx context.Context, tx *sql.Tx, orderCode string) error {
	query := "INSERT INTO `subscription` (`order_code`, `id`) VALUES (?, ?)"
	_, err := tx.ExecContext(ctx, query, orderCode, s.SubscriptionID)

	return err
}
