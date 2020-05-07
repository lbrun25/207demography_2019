package demography

import (
	"fmt"
	"strings"
)

// Years = slice of year
var Years []int

// Demographies - slice of Demography
var Demographies []Demography

// Demography - Struct which holds demography data
type Demography struct {
	countryName string
	countryCode string
	values []int
}

// Populations - slice of population by year
var Populations []Population

// Population - struct which holds population's infos
type Population struct {
	year int
	value int
}

func getCountryNames() []string {
	var countryNames []string
	for _, demography := range Demographies {
		countryNames = append(countryNames, demography.countryName)
	}
	return countryNames
}

func displayResults() {
	fmt.Println("Country:", strings.Join(getCountryNames(), ", "))

	fmt.Println("Fit1")
	fmt.Printf("\tY= %.2f X - %.2f\n", 0.0, 0.0)
	fmt.Printf("\tRoot-mean-square deviation: %.2f\n", 0.0)
	fmt.Printf("\tPopulation in 2050: %.2f\n", 0.0)

	fmt.Println("Fit2")
	fmt.Printf("\tY= %.2f Y + %.2f\n", 0.0, 0.0)
	fmt.Printf("\tRoot-mean-square deviation: %.2f\n", 0.0)
	fmt.Printf("\tPopulation in 2050: %.2f\n", 0.0)

	fmt.Println("Correlation", 0.0)
}

// Main - Demography main
func Main() {
	ComputeYearMean()
	ComputeMeanValues()
	ComputePopulations()
	ComputeFit1()
	displayResults()
}