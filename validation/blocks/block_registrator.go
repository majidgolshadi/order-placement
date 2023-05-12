package blocks

const (
	productItemSectionName = "product_items"
	paymentSectionName     = "payment"
)

type OrderSkeleton struct {
	ProductItems []map[string]any `json:"product_items"`
	Payment      Payment          `json:"payment"`
}

type BlockType struct {
	Type string `json:"type"`
}

type registrar struct {
	productsItems map[string]DataUnmarshal
}

func GetBlockRepository() *registrar {
	r := &registrar{}
	r.registerProductItemTypes()

	return r
}

func (r *registrar) registerProductItemTypes() {
	r.productsItems = make(map[string]DataUnmarshal)

	r.productsItems["food"] = Food{}
	r.productsItems["grocery"] = Grocery{}
}

func (r *registrar) GetProductStruct(productType string, data map[string]any) SectionValidator {
	sv, _ := r.productsItems[productType].MapUnmarshal(data)
	return sv
}
