package main

import (
	"exportedapp/staff"
	"log"
)

var employees = []staff.Employee{
	{FirstName: "John", LastName: "Smith", Salary: 50000, FullTime: false},
	{FirstName: "Sally", LastName: "Johns", Salary: 60000, FullTime: true},
	{FirstName: "Mark", LastName: "Smith", Salary: 15000, FullTime: false},
	{FirstName: "Margaret", LastName: "Curter", Salary: 100000, FullTime: false},
}

func main() {
	myStaff := staff.Office{
		AllStaff: employees,
	}

	log.Println(myStaff.All())
	log.Println("Overpaid staff:", myStaff.Overpaid())

	staff.OverPaidLimit = 50000
	log.Println("Overpaid staff:", myStaff.Overpaid())
	log.Println("Underpaid staff: ", myStaff.Underpaid())
}
