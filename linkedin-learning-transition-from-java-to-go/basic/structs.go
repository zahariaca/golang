package main

import "fmt"

type city struct {
	name  string
	tempC float64
}

func newCity(n string) city {
	return city{
		name: n,
	}
}

func (c city) tempF() float64 {
	return c.tempC*1.8 + 32
}

func main() {
	c := newCity("San Francisco")
	c.tempC = 17.0
	tempF := c.tempF()

	fmt.Printf("[%v]:tempC=%v;tempF=%v \n", c.name, c.tempC, tempF)
}
