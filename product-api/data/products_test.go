package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "Nicks",
		Price: 1.23,
		SKU:   "abs-abc-defs",
	}
	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
