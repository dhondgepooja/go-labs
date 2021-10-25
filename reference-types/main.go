package main

import (
	"fmt"
	"github.com/eiannone/keyboard"
	"log"
	"sort"
)

type Animal struct {
	Name         string
	Sound        string
	NumberOfLegs int
}

func (a *Animal) Says() {
	fmt.Printf("A %s says %s", a.Name, a.Sound)
}

var keyPressChan chan rune // define a channel that takes a rune. Channel can take in only a single arg

func main() {
	//Pointers
	//studyPointers()

	// Slices
	//studySlices()

	// Maps
	//studyMaps()

	// Functions
	//studyFunctions()

	// Channels
	//studyChannels()

	// Interfaces
	studyInterfaces()
}

type Dog struct {
	Name         string
	Sound        string
	NumberOfLegs int
}

type Cat struct {
	Name         string
	Sound        string
	NumberOfLegs int
	HasTail      bool
}

type AnimalInter interface {
	Says() string
	HowManyLegs() int
}

func Riddle(a AnimalInter) {
	riddle := fmt.Sprintf(`This animal says "%s" and has "%d" legs`, a.Says(), a.HowManyLegs())
	fmt.Println(riddle)
}

func (d *Dog) Says() string {
	return d.Sound
}

func (d *Dog) HowManyLegs() int {
	return d.NumberOfLegs
}

func (c *Cat) Says() string {
	return c.Sound
}

func (c *Cat) HowManyLegs() int {
	return c.NumberOfLegs
}

func studyInterfaces() {
	dog := Dog{
		Name:         "Dog",
		Sound:        "woof",
		NumberOfLegs: 4,
	}
	Riddle(&dog)

	cat := Cat{
		Name:         "cat",
		Sound:        "meow",
		NumberOfLegs: 4,
	}

	Riddle(&cat)
}

func studyChannels() {
	keyPressChan = make(chan rune)

	go listenForKeyPress()

	fmt.Println("Press any key, or q to quit")

	_ = keyboard.Open()

	defer func() {
		keyboard.Close()
	}()

	for {
		char, _, _ := keyboard.GetSingleKey()
		if char == 'q' || char == 'Q' {
			break
		}

		keyPressChan <- char
	}
}

func listenForKeyPress() {
	for {
		key := <-keyPressChan
		fmt.Println("You pressed", string(key))
	}
}

func studyFunctions() {
	fmt.Println("Sum is:", addTwoNumbers(3, 4))
	fmt.Println("Sum of many is:", sumMany(1, 2, 3, 4))

	var dog Animal
	dog.Name = "dog"
	dog.Sound = "woof"

	dog.Says()

	cat := Animal{
		Name:         "cat",
		Sound:        "meow",
		NumberOfLegs: 4,
	}

	cat.Says()
}

func studyMaps() {
	intMap := make(map[string]int)

	intMap["One"] = 1
	intMap["Two"] = 2
	intMap["Three"] = 3

	for key, val := range intMap {
		fmt.Println(key, val) // doesn't follow any order
	}

	delete(intMap, "Three")

	for key, val := range intMap {
		fmt.Println(key, val) // doesn't follow any order
	}

	el, ok := intMap["Two"]
	if ok {
		fmt.Println(el, "is in map")
	} else {
		fmt.Println(el, "is not in map")
	}
}

func studySlices() {
	var animals []string
	animals = append(animals, "dog")
	animals = append(animals, "fish")
	animals = append(animals, "cat")
	animals = append(animals, "horse")

	log.Println(animals)

	for _, animal := range animals {
		fmt.Println(animal)
	}

	for i, animal := range animals {
		fmt.Println(i, animal)
	}

	fmt.Println("First two elements are", animals[0:2])
	fmt.Println("The slice is", len(animals), "elements long")

	fmt.Println("Is it sorted?", sort.StringsAreSorted(animals))
	sort.Strings(animals)
	fmt.Println("Is it sorted?", sort.StringsAreSorted(animals))
	fmt.Println(animals)

	animals = DeleteFromSlice(animals, 1)
	fmt.Println(animals) // NOTE that this broke our sorting
	sort.Strings(animals)
	fmt.Println(animals)
}

func studyPointers() {
	x := 10
	myFirstPointer := &x

	fmt.Println("x is ", x)
	fmt.Println("myFirstPointer is", myFirstPointer)

	*myFirstPointer = 15
	fmt.Println("x is now ", x)
	fmt.Println("myFirstPointer is", myFirstPointer)

	changeValueOfPointer(&x)
	fmt.Println("After func call x is now ", x)
}

func changeValueOfPointer(num *int) {
	*num = 25
}

func DeleteFromSlice(a []string, i int) []string {
	a[i] = a[len(a)-1]
	a[len(a)-1] = ""
	a = a[:len(a)-1]
	return a
}

func addTwoNumbers(x, y int) int {
	// sum = x + y
	//return // NOTE: naked return shou;d't be used often
	return x + y
}

func sumMany(nums ...int) int {
	total := 0

	for _, x := range nums {
		total = total + x
	}

	return total
}
