package blocks

type SectionValidator interface {
	IsBlockDataValid() bool
	IsPaymentSectionValid(payment Payment) bool
}

type DataUnmarshal interface {
	MapUnmarshal(map[string]any) (SectionValidator, error)
}
