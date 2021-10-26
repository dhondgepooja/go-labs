package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println()

	name := "Pooja"
	fmt.Println("String:", name)
	fmt.Println()

	fmt.Println("Bytes")
	for i := 0; i < len(name); i++ {
		fmt.Printf("%x ", name[i])
	}

	fmt.Println()
	fmt.Println()

	fmt.Println("Index\tRune\tString")
	for x, y := range name {
		fmt.Println(x, "\t", y, "\t", string(y))
	}

	fmt.Println()
	fmt.Println("Three ways to concatenate")
	h := "Hello, "
	w := "world."

	myString := h + w
	fmt.Println("myString: ", myString)

	// using fmt
	myString = fmt.Sprintf("%s%s", h, w) // efficient than above
	fmt.Println("myString: ", myString)

	// string builder
	var sb strings.Builder
	sb.WriteString(h)
	sb.WriteString(w)
	myString = sb.String()
	fmt.Println("myString: ", myString)

	// substring
	name = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	fmt.Println(name[10:13])

	courseName := "Go: Learn Go for Begineers Crash Course"
	fmt.Println(string(courseName[6]))

	for i := 13; i <= 21; i++ {
		fmt.Print(string(courseName[i]))
	}

	fmt.Println("Length:", len(courseName))

	var mySlice []string
	mySlice = append(mySlice, "hello")
	mySlice = append(mySlice, "world")

	fmt.Println("Last element in slice:", mySlice[len(mySlice)-1])

	courses := []string{
		"Learn Go for Begineers",
		"Learn Java for Begineers",
		"Learn Ruby for Begineers",
	}

	for i, x := range courses {
		if strings.Contains(x, "Go") {
			fmt.Println("Go found at", i)
		}
	}

	if strings.Contains(courseName, "Go") {
		courseName = strings.Replace(courseName, "Go", "GoLang", 1)
		fmt.Println("Course name: ", courseName)
	}
}
