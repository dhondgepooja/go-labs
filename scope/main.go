package main

import (
	"fmt"
	"myscopeapp/packageone"
)

var one = "One" // package level variable
var myVar = "myVar"

func main() {
	var one = "main one" // this is called variable shadowing BAD PRACTICE
	fmt.Println(one)
	myFunc()

	newString := packageone.PublicVar
	fmt.Println(newString)

	packageone.Exported()

	//assignment
	var blockVar = "blockVar"
	packageone.PrintMe(myVar, blockVar)
}

func myFunc() {
	fmt.Println(one)
}
