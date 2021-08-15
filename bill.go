package main

import (
	"fmt"
)

type bill struct {
	name string
	items map[string] float64
	tip float64
}

func newBill(name string) bill {
	b:= bill{
		name:  name,
		items: map[string]float64{"pie":5.99,"cake":3.99},
		tip:   100,
	}
	return b
}
//function to format bill.
//We associate the format function with the bill through the receiver.
//the receiver below(b) has limited the function to only be able to format the bill object
func (b bill) format() string {
fs:="Bill breakdown\n"
	var total float64= 0
	//list items
	for k,v:=range b.items{
		fs+=fmt.Sprintf("%-25v ...$%v \n", k+":", v)
		total+=v
	}
	//add tip
	fs+=fmt.Sprintf("%-25v ...$%0.2f", "total:", b.tip)

	//total
	fs+=fmt.Sprintf("%-25v ...$%0.2f", "total:", total+b.tip)
	return fs
}

//update tip
func (b *bill) updateTip(tip float64) {
	b.tip=tip
}
// add tip
func (b *bill) addItem(name string, price float64)   {
	b.items[name]=price
}

