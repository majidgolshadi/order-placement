package blocks

type Payment struct {
	TotalOrderValue  string `json:"total_order_value"`
	FoodItemPrice    string `json:"food_item_price"`
	GroceryItemPrice string `json:"grocery_item_price"`
}

func (p Payment) IsBlockDataValid() bool {
	if p.TotalOrderValue == "" {
		return false
	}

	return true
}
