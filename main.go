package main

import "fmt"

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
}
