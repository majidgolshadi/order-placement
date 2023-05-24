package internal

import (
	"errors"
	"fmt"
	"time"
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
			t := time.Now()
			return fmt.Sprintf("%s-%s-%d%d", prefix, "12313", t.Year(), t.Month()), nil
		}
	}

	return "", errors.New("order type does not found")
}
