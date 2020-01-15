package main

import (
	"fmt"
	"time"
)

// Employee ...
type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

func main() {
	var empty Employee
	fmt.Println("empty: ", empty)

	dilbert := Employee{
		ID:       1,
		Name:     "Dilbert",
		Position: "Engineer",
		Salary:   5000,
	}
	fmt.Println("dilbert:", dilbert)

	dilbert.Salary -= 5000 // demoted, for writing too few lines of code

	position := &dilbert.Position
	*position = "Senior " + *position // promoted, for outsourcing to Elbonia

	fmt.Println("dilbert:", dilbert)

	alice := &Employee{
		ID:   2,
		Name: "Alice",
	}
	fmt.Println("alice:", alice)

	fmt.Println(alice.ID, alice.Name)
}
