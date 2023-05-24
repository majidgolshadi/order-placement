package blocks

import "errors"

type Payment struct {
	TotalOrderValue  string `json:"total_order_value"`
	FoodItemPrice    string `json:"food_item_price"`
	GroceryItemPrice string `json:"grocery_item_price"`
}

func (p Payment) IsBlockDataValid() error {
	if p.TotalOrderValue == "" {
		return errors.New("total order value is empty")
	}

	return nil
}
