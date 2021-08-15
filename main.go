package main

import "fmt"

func main() {
	mybill:= newBill("IceCream Bill")
	fmt.Println(mybill.format())
}
