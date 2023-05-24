package internal

import (
	"errors"
	"fmt"
)

type codeGenerator struct {
	prefix map[string]string
}

func NewCodeGenerator() *codeGenerator {
	return &codeGenerator{
		prefix: map[string]string{
			"food":         "fo",
			"din-in":       "pro",
			"subscription": "sub",
		},
	}
}

func (cg *codeGenerator) GetOrderCode(orderType string) (string, error) {
	for key, prefix := range cg.prefix {
		if orderType == key {
			return fmt.Sprintf("%s-%s", prefix, "12313"), nil
		}
	}

	return "", errors.New("order type does not found")
}
