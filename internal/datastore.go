package internal

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/deliveryhero/pd-order-placement/internal/blocks"
	"github.com/mitchellh/mapstructure"
)

type Repository interface {
	// Persist method according to data store technology can be defined differently
	Persist(ctx context.Context, tx *sql.Tx, orderCode string) error
}

type Datastore struct {
	db *sql.DB
}

func (ds *Datastore) newConnection() error {
	// build the DSN
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", "root", "root", "127.0.0.1", 3306, "test_db")
	// Open the connection
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	ds.db = db
	return nil
}

func (ds *Datastore) Flush(ctx context.Context, payload payloadSkeleton, orderCode string) error {
	// TODO: error handling
	_ = ds.newConnection()

	// begin Transaction
	tx, err := ds.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if pbErr := ds.PersistProductItemsSection(ctx, tx, orderCode, payload); pbErr != nil {
		return pbErr
	}

	if cErr := tx.Commit(); cErr != nil {
		return cErr
	}

	return nil
}

func (ds *Datastore) PersistProductItemsSection(ctx context.Context, tx *sql.Tx, orderCode string, payload payloadSkeleton) error {
	for _, productItem := range payload.ProductItems {
		r, err := ds.getProductItemBlockRepository(productItem["type"].(string), productItem)
		if err != nil {
			return err
		}

		if pErr := r.Persist(ctx, tx, orderCode); pErr != nil {
			return pErr
		}

	}

	return nil
}

func (ds *Datastore) getProductItemBlockRepository(productItemType string, data map[string]any) (Repository, error) {
	if productItemType == "food" {
		var productItem blocks.ProductItemFood
		if err := mapstructure.Decode(data, &productItem); err != nil {
			return nil, err
		}

		return productItem, nil
	}

	if productItemType == "grocery" {
		var productItem blocks.ProductItemGrocery
		if err := mapstructure.Decode(data, &productItem); err != nil {
			return nil, err
		}

		return productItem, nil
	}

	if productItemType == "subscription" {
		var productItem blocks.ProductItemSubscription
		if err := mapstructure.Decode(data, &productItem); err != nil {
			return nil, err
		}

		return productItem, nil
	}

	return nil, errors.New("unknown type")
}
