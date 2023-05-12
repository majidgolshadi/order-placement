package blocks

import (
	"github.com/mitchellh/mapstructure"
)

type Food struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (f Food) MapUnmarshal(data map[string]any) (SectionValidator, error) {
	var food Food
	err := mapstructure.Decode(data, &food)

	return food, err
}
func (f Food) IsBlockDataValid() bool {
	if f.ID == "" {
		return false
	}

	if f.Name == "" {
		return false
	}

	return true
}

func (f Food) IsPaymentSectionValid(payment Payment) bool {
	if payment.FoodItemPrice == "" {
		return false
	}

	return true
}
