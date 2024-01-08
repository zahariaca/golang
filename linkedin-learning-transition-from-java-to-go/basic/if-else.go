package main

import "fmt"

func printParity(x int) {
	if x%2 == 0 {
		fmt.Printf("%v is even\n", x)
		return
	}

	fmt.Printf("%v is odd\n", x)
}

func printParityWithShorthandVariableDeclaration(x int) {
	if r := x % 2; r == 0 {
		fmt.Printf("r has the value: %v\n", r)
		fmt.Printf("%v is even\n", x)
		return
	}

	// r is not in scope here
	fmt.Printf("%v is odd\n", x)
}

func main() {
	printParity(4)
	printParity(5)
	printParityWithShorthandVariableDeclaration(4)
	printParityWithShorthandVariableDeclaration(5)
}
