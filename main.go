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

	fmt.Printf("hello %v, your age is %v \n", name, age)
	fmt.Printf("hello %q, your age is %q \n", name, age)
	fmt.Printf("hello %T, your age is %T \n", name, age)
}
