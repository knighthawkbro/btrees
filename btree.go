package main

import (
	"btrees/treemap"
	"fmt"
)

// Map interface
type Map interface {
	Add(key string, value int) int
	Remove(key string) int
	Get(key string) int
	Contains(key string) bool
	Size() int
	String() string
}

func main() {
	tree := treemap.New()
	driver(tree)
	tree = treemap.New()
	seuss(tree)
}

func driver(ages Map) {
	names := []string{
		"Ken", "Fred", "Sam", "Ben", "Harold",
		"Tom", "James", "Paul", "Jeff", "George", "Jan",
	}
	numbers := []int{
		27, 21, 22, 21, 16,
		30, 28, 29, 24, 25, 25,
	}

	for index, name := range names {
		ages.Add(name, numbers[index])
	}

	// Add
	fmt.Println("Here's the group with each person's age\n", ages)

	// Contains
	fmt.Print("Is George in the group? ")
	if ages.Contains("George") {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
	fmt.Print("Is Keith in the group? ")
	if ages.Contains("Keith") {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}

	// Get
	fmt.Printf("How old is Harold? ")
	age := ages.Get("Harold")
	if age >= 0 {
		fmt.Println(age)
	} else {
		fmt.Println("<nil>")
	}
	fmt.Printf("How old is Lily? ")
	age = ages.Get("Lily")
	if age >= 0 {
		fmt.Println(age)
	} else {
		fmt.Println("<nil>")
	}
}

func seuss(words Map) {

}
