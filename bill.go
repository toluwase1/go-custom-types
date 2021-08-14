package main

type bill struct {
	name string
	items map[string] float64
	tip float64
}

func NewBill(name string) bill {
	b:= bill{
		name:  "Tolu",
		items: map[string]float64{"abc":1,"web":5},
		tip:   100,
	}
	return b
}

