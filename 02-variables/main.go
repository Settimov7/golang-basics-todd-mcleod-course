package main

import "fmt"

type notANumber int

var globalString = "azaza"
var zeroValueString string

func main() {
	number := 33

	fmt.Printf("%T\n", number)
	fmt.Printf("%v\n", number)
	fmt.Printf("%#v\n", number)
	fmt.Printf("%b\n", number)
	fmt.Printf("%x\n", number)
	fmt.Printf("%#x\n", number)

	result := fmt.Sprintf("%b\t%x\t%#x", number, number, number)

	fmt.Println("result", result)

	var someValue notANumber = 666

	fmt.Printf("someValueType %T\n", someValue)

	printString()

	printZeroValueString()
	changeZeroValueString()
	printZeroValueString()
}

func printString() {
	fmt.Println(globalString)
}

func changeZeroValueString() {
	zeroValueString = `
		not
		empty
		string
	`
}

func printZeroValueString() {
	fmt.Println("zeroValueString =", zeroValueString)
}
