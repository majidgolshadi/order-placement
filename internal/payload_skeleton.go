package internal

import (
	"encoding/json"
	"github.com/deliveryhero/pd-order-placement/internal/blocks"
)

type payloadSkeleton struct {
	Type         string           `json:"type"`
	ProductItems []map[string]any `json:"product_items"`
	Payment      blocks.Payment   `json:"payment"`
}

func GetPayloadSkeleton(payload []byte) (payloadStr payloadSkeleton, err error) {
	err = json.Unmarshal(payload, &payloadStr)
	return
}
