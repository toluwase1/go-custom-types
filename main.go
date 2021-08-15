package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	mybill:= newBill("New Bill")
	//mybill.updateTip(20)
	//mybill.addItem("biscuit", 30)
	promptOptions(mybill)
	//fmt.Println(mybill.format())
//createBill()

}

func createBill() bill  {
	reader:=bufio.NewReader(os.Stdin)
	//fmt.Print("Create a new bill name")
	//name,_:=reader.ReadString('\n')
	//name = strings.TrimSpace(name)
	name,_ := getInput("Create a new bill name", reader)
	b:= newBill(name)
	fmt.Println("Created the bill", b.name)
	return b
}

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	return strings.TrimSpace(input), err
}

func promptOptions(b bill)  {
	reader :=bufio.NewReader(os.Stdin)
	opt,_ := getInput("Choose option: (a - add item, s - save, t-add tip): ", reader)

	switch opt {
	case "a":
		name, _ := getInput("Item Name: ", reader)
		price,_ := getInput("Item Price: ", reader)
		p,err:=strconv.ParseFloat(price,64)
		if err != nil {
			fmt.Print("The price must be a number")
			promptOptions(b)
		}
		b.addItem(name, p)
		fmt.Println("Items added ", name, p)
		promptOptions(b)
	case "t":
		tip, _:= getInput("Enter tip amount in $", reader)

		t,err:=strconv.ParseFloat(tip,64)
		if err != nil {
			fmt.Print("The tip must be a number <3 ")
			promptOptions(b)
		}
		b.updateTip(t)
		fmt.Println("Tip value updated: ", t)
		promptOptions(b)
	case "s":
		b.save()
		fmt.Println("You saved the file: ", b.name)
	default:
		fmt.Println("Your option is invalid...")
		promptOptions(b)
	}
}

//save bill
func (b *bill) save()  {
	data:= [] byte (b.format()) //format and put into a byteslice, then save into data
	err:= os.WriteFile("bills/"+b.name+".txt", data,0666)
	if err != nil {
		panic(err)
	}
	fmt.Print("Bill was saved to the file")

}
