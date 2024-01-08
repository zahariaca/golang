package main

import "fmt"

// we use a *int as the return type to indicate, in case of an error,
// that the return value is unsuable (otherwise we would have returned a 0)
func area(x int, y int) (*int, error) {
	if x == 0 || y == 0 {
		return nil, fmt.Errorf("zero area: [%v,%v]", x, y)
	}

	area := x * y
	return &area, nil
}

func main() {
	x := 2
	y := 0

	area, err := area(x, y)
	if	err != nil {
		fmt.Println("error: ", err)
		return
	}

	fmt.Printf("area(%v,%v) = %v \n", x, y, *area)
}
