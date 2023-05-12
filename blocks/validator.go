package blocks

import (
	"encoding/json"
	"errors"
	"github.com/mitchellh/mapstructure"
)

type BlockValidator interface {
	IsBlockDataValid() error
	IsPaymentSectionValid(payment Payment) error
}

type orderPayloadSkeleton struct {
	ProductItems []map[string]any `json:"product_items"`
	Payment      Payment          `json:"payment"`
}

type Validator struct {
}

func (v *Validator) IsDataValid(payload []byte) error {
	payloadStr := orderPayloadSkeleton{}
	if err := json.Unmarshal(payload, &payloadStr); err != nil {
		return err
	}

	if err := payloadStr.Payment.IsBlockDataValid(); err != nil {
		return err
	}

	for _, productItem := range payloadStr.ProductItems {
		bv, err := v.getProductItemBlockValidator(productItem["type"].(string), productItem)
		if err != nil {
			return err
		}

		if bErr := bv.IsBlockDataValid(); bErr != nil {
			return bErr
		}

		if pErr := bv.IsPaymentSectionValid(payloadStr.Payment); pErr != nil {
			return pErr
		}
	}

	return nil
}

func (v *Validator) getProductItemBlockValidator(productItemType string, data map[string]any) (BlockValidator, error) {
	if productItemType == "food" {
		var productItem ProductItemFood
		if err := mapstructure.Decode(data, &productItem); err != nil {
			return nil, err
		}

		return productItem, nil
	}

	if productItemType == "grocery" {
		var productItem ProductItemGrocery
		if err := mapstructure.Decode(data, &productItem); err != nil {
			return nil, err
		}

		return productItem, nil
	}

	return nil, errors.New("unknown type")
}
