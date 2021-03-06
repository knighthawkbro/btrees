package main

import (
	"btrees/set"
	"btrees/treemap"
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
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

type Set interface {
	Add(item string) bool
	Remove() string
	Get() string
	Contains(item string) bool
	Size() int
	String() string
}

func main() {
	fmt.Println("\n**************************************************")
	fmt.Print("*\tRunning Tree Map driver function...")
	fmt.Println("\n**************************************************")
	fmt.Println("")
	tree := treemap.New()
	driver(tree)
	fmt.Println("\n**************************************************")
	fmt.Print("*\tRunning Tree Map suess function...")
	fmt.Println("\n**************************************************")
	fmt.Println("")
	tree = treemap.New()
	suess(tree)
	fmt.Println("\n**************************************************")
	fmt.Print("*\tRunning Tree Map suess function...")
	fmt.Println("\n**************************************************")
	fmt.Println("")
	treeset := set.New()
	seussSet(treeset)
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

func suess(words Map) {
	file, err := os.Open("greenEggs.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	reg, err := regexp.Compile("[^a-zA-Z]+")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		for _, word := range strings.Split(scanner.Text(), " ") {
			word = reg.ReplaceAllString(word, "")
			if word == "" {
				continue
			}
			if words.Contains(strings.ToLower(word)) {
				count := words.Get(strings.ToLower(word))
				if count < 0 {
					continue
				}
				words.Add(strings.ToLower(word), count+1)
				continue
			}
			words.Add(strings.ToLower(word), 1)
		}
	}
	fmt.Println(words)
}

func seussSet(words Set) {
	file, err := os.Open("greenEggs.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	reg, err := regexp.Compile("[^a-zA-Z]+")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		for _, word := range strings.Split(scanner.Text(), " ") {
			word = reg.ReplaceAllString(word, "")
			words.Add(strings.ToLower(word))
		}
	}
	fmt.Println(words)
}
