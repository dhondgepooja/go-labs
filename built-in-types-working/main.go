package main

import (
	"fmt"
	"log"
)

var myInt int // most commonly used
// Supports this so that if you want to use the app on phone or specific devices that only support 32 bit or 64 bit. This is discouraged
//var myInt16 int16 // almost never use
//var myInt32 int32
//var myInt64 int64

var myUint uint

var myFloat float32 // this will do the job in most cases
var myFloat64 float64

// Structs

type Car struct {
	NumberOfTypes int
	Luxury        bool
	Make          string
	Model         string
	Year          int
}

func main() {
	myInt = -10
	myUint = 10
	myFloat = 10.1
	myFloat64 = 100.5

	log.Println(myInt, myUint, myFloat, myFloat64)

	myString := "Pooja"

	log.Println(myString)

	myString = "Pooja2" // creates new string as strings are immutable

	var myBool = true
	log.Println(myBool)

	// Aggregate Types:

	// Arrays: These are not generally used. Instead use slices
	var myStrings [3]string
	myStrings[0] = "cat"
	myStrings[1] = "dog"
	myStrings[2] = "fish"

	fmt.Println("First elements: ", myStrings[0])

	// Structs:
	var myCar Car

	myCar.Make = "Tesla"
	myCar.Model = "model 3"
	myCar.Luxury = true
	myCar.Year = 2019

	myCar2 := Car{
		Make:   "Tesla",
		Model:  "Model3",
		Luxury: true,
		Year:   2019,
	}

	log.Println(myCar, myCar2)

}
