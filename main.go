package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

	var name string
	var phone int
	var amount float64

	phone = 81238484
	amount = 33.2

	age := 2
	fmt.Println(name, age, phone, amount)
	fmt.Print("hello \n")

	// formatted string
	fmt.Printf("hello %v, your age is %v \n", name, age)
	fmt.Printf("hello %q, your age is %q \n", name, age)
	fmt.Printf("hello %T, your age is %T \n", name, age)

	// array
	var ages [3]int = [3]int{23, 4, 5}
	var agesTwo = [3]int{1, 2, 3}

	fmt.Println(ages, len(ages))
	fmt.Println(agesTwo, len(agesTwo))

	var names [3]string = [3]string{"ayo", "david", "dav"}

	fmt.Println(names, len(names))

	// slices
	var words = []string{"he", "she", "yes"}
	fmt.Println(words)
	// append
	words = append(words, "goat")
	fmt.Println(words)
	// rang
	fmt.Println(words[1:3])

	// string package
	sentence := "Hello i am a boy"
	fmt.Println(strings.Contains(sentence, "Hello"))
	fmt.Println(strings.ToUpper(sentence))

	// sort
	scores := []int{1, 34, 55, 354, 43, 3}
	sort.Ints(scores)
	fmt.Println(scores)
	fmt.Println(sort.SearchInts(scores, 1))

	people := []string{"shina", "tomiwa", "dami", "kator", "peace"}
	sort.Strings(people)
	fmt.Println(people)
	fmt.Println(sort.SearchStrings(people, "peace"))

	// loops
	// while loop in go
	x := 0
	for x < 5 {
		fmt.Println("value of x is", x)
		x++
	}

	// similar to for loop in go
	for i := 0; i < 5; i++ {
		fmt.Println("value of x is", i)
	}

	// loop through a slices
	for i := 0; i < len(people); i++ {
		fmt.Println(people[i])
	}

	for index, value := range people {
		fmt.Printf("the value at index %v is %v \n", index, value)
	}

}
