package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "Nic",
		Price: 1.00,
		SKU:   "abs-asd-asd",
	}
	err := p.Validate()
	if err != nil {
		t.Fatal(err)
	}
}
