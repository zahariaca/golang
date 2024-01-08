package main

import "fmt"

var (
	v0 int
	v1 float64
	v3 bool
	v4 string
)

func main() {
	fmt.Printf("[v0]: variable type = %T, value = %v\n", v0, v0)
	fmt.Printf("[v1]: variable type = %T, value = %v\n", v1, v1)
	fmt.Printf("[v3]: variable type = %T, value = %v\n", v3, v3)
	fmt.Printf("[v4]: variable type = %T, value = %v\n", v4, v4)
}
