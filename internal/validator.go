package internal

import (
	"errors"
	"github.com/deliveryhero/pd-order-placement/internal/blocks"
	"github.com/mitchellh/mapstructure"
)

type BlockValidator interface {
	IsBlockDataValid() error
}

type PaymentCrossDependencyValidator interface {
	IsPaymentSectionValid(payment blocks.Payment) error
}

type ProductBlockValidator interface {
	BlockValidator
	PaymentCrossDependencyValidator
}

type Validator struct {
}

func (v *Validator) IsDataValid(payload payloadSkeleton) error {
	if err := payload.Payment.IsBlockDataValid(); err != nil {
		return err
	}

	return v.IsProductItemsSectionValid(payload)
}

func (v *Validator) IsProductItemsSectionValid(payload payloadSkeleton) error {
	for _, productItem := range payload.ProductItems {
		bv, err := v.getProductItemBlockValidator(productItem["type"].(string), productItem)
		if err != nil {
			return err
		}

		if bErr := bv.IsBlockDataValid(); bErr != nil {
			return bErr
		}

		if pErr := bv.IsPaymentSectionValid(payload.Payment); pErr != nil {
			return pErr
		}
	}

	return nil
}

func (v *Validator) getProductItemBlockValidator(productItemType string, data map[string]any) (ProductBlockValidator, error) {
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
