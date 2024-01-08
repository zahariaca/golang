package main

import "fmt"

func main() {
	var ptr *string
	greeting := "Hello, LinkedIn!"

	ptr = &greeting

	fmt.Println("Greeting: ", greeting)
	fmt.Println("Address of greeting: ", ptr)
	fmt.Println("Value stored in ptr: ", *ptr)
}
