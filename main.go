package main

import "fmt"

func main() {
	var name string = "Go Microservices"
	age := 31
	const pi = 3.14
	fmt.Println("Hello, Go!", name, age, pi)
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	if age >= 18 {
		fmt.Println("Adult")
	} else {
		fmt.Println("Minor")
	}
}
