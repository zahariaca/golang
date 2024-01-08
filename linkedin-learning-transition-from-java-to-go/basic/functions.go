package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func negatePointer(x *int) {
	neg := -*x
	*x = neg
}

func negateNonPointer(x int) {
	neg := -x
	x = neg
}

func rectangle(x int, y int) (int, int) {
	area := x * y
	circumf := (x + y) * 2
	return area, circumf
}

func main() {
	x := 3
	y := 2

	sub := func(x int, y int) int {
		return x - y
	}

	fmt.Printf("add(%v, %v): %v\n", x, y, add(x, y))
	fmt.Printf("sub(%v, %v): %v\n", x, y, sub(x, y))
	negateNonPointer(x)
	fmt.Printf("expect value to not change: %v\n", x)
	negatePointer(&x)
	fmt.Printf("expect value to change: %v\n", x)
	negatePointer(&x)

	area, circumf := rectangle(x, y)
	fmt.Printf("rectangle(%v, %v): area: %v\n", x, y, area)
	fmt.Printf("rectangle(%v, %v): circumf: %v\n", x, y, circumf)
}
