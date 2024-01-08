package main

import (
	"fmt"

	"temperature-service/models"
	"temperature-service/printer"
)

func main() {
	fmt.Printf("Welcome to the temperature service!\n\n")

	// initialise printer and defer cleanup
	p := printer.New()
	defer p.Cleanup()

	p.CityHeader()

	// setup 3 cities

	lon := models.NewCity("London", 23, false, false)
	bcn := models.NewCity("Barcelona", 30, true, false)
	nyc := models.NewCity("New York", 28, true, false)
	ant := models.NewCity("St Anton", -3, false, true)
	asp := models.NewCity("Aspen", -5, false, true)

	// print city details
	p.CityDetails(lon)
	p.CityDetails(bcn)
	p.CityDetails(nyc)
	p.CityDetails(ant)
	p.CityDetails(asp)
}
