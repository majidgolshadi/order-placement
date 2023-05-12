package blocks

type SectionValidator interface {
	IsBlockDataValid() error
	IsPaymentSectionValid(payment Payment) error
}

type DataUnmarshal interface {
	MapUnmarshal(map[string]any) (SectionValidator, error)
}
