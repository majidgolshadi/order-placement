package blocks

import (
	"errors"
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
func (f Food) IsBlockDataValid() error {
	if f.ID == "" {
		return errors.New("id field is empty")
	}

	if f.Name == "" {
		return errors.New("name field is empty")
	}

	return nil
}

func (f Food) IsPaymentSectionValid(payment Payment) error {
	if payment.FoodItemPrice == "" {
		return errors.New("food item price should be set")
	}

	return nil
}
