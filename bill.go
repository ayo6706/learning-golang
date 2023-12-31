package main

import "fmt"

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{"pie": 35.7, "cake": 55.8},
		tip:   0,
	}

	return b
}

func (b bill) format() string {
	fs := "Bill break: \n"
	var total float64 = 0

	for k, v := range b.items {
		fs += fmt.Sprintf("%v ...$%v \n", k+":", v)
		total += v
	}

	fs += fmt.Sprintf("%-25v ...$%0.2f", "total:", total)
	return fs
}
