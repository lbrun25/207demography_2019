package demography

import (
	"math"
)

const (
	targetYear = 2050
	fit1 = 0
	fit2 = 1
)

// MeanYear - mean of years
var MeanYear float64

// MeanYear - mean of values
var MeanValues float64

type fitValues struct {
	a float64
	b float64
	rootMeanSquareDeviation float64
	population float64
}

// Fits - struct which holds values for both fit
type Fits struct {
	one fitValues
	two fitValues
}

// Fit - Fit's holder
var Fit Fits

// ComputeYearMean - get mean date (year)
func ComputeYearMean() {
	sum := 0.0

	for i := 0; i < len(Years); i++ {
		sum += float64(Years[i])
	}
	MeanYear = sum /  float64(len(Years))
}

func getA(fit int) float64 {
	numerator := 0.0
	denominator := 0.0

	for _, population := range Populations {
		switch fit {
		case fit1:
			numerator += (float64(population.value) - MeanValues) * (float64(population.year) - MeanYear)
			denominator += math.Pow(float64(population.year) - MeanYear, 2)
		case fit2:
			numerator += (float64(population.year) - MeanYear) * (float64(population.value) - MeanValues)
			denominator += math.Pow(float64(population.value) - MeanValues, 2)
		}
	}
	return numerator / denominator
}

func getRootMeanStandardDeviation(fit int, a float64, b float64) float64 {
	numerator := 0.0
	y := 0.0

	for i, year := range Years {
		for _, demography := range Demographies {
			y += float64(demography.values[i])
		}
		switch fit {
		case fit1:
			numerator += math.Pow((a * float64(year) + b) - y, 2)
		case fit2:
			numerator += math.Pow(((float64(year) - b) / a) - y, 2)
		}
		y = 0
	}
	return math.Sqrt(numerator / float64(len(Years)))
}

// ComputeFit1 - get fit1
func ComputeFit1() {
	a := getA(fit1)
	b := MeanValues - a * MeanYear
	population := a * targetYear + b

	Fit = Fits{one: fitValues{
		a: a,
		b: b,
		rootMeanSquareDeviation: getRootMeanStandardDeviation(fit1, a, b),
		population: population,
	}}
}

// ComputePopulations - fill Populations's slice
func ComputePopulations() {
	y := 0

	for i, year := range Years {
		for _, demography := range Demographies {
			y += demography.values[i]
		}
		Populations = append(Populations, Population{year: year, value: y})
		y = 0
	}
}

// ComputeMeanValues - get mean demography values
func ComputeMeanValues() {
	sum := 0.0

	for _, demography := range Demographies {
		for _, value := range demography.values {
			sum += float64(value)
		}
	}
	MeanValues = sum / float64(len(Years))
}

// GetCoefficients - Get linear coefficients
func GetCoefficients() float64 {
	res := 0.0
	return res
}

// GetRootMeanSquareDeviation - Get root-mean-square deviation
func GetRootMeanSquareDeviation() float64 {
	res := 0.0
	return res
}

// GetPopulationPrediction - Get population prediction in 2050
func GetPopulationPrediction() float64 {
	res := 0.0
	return res
}

// GetCorrelation - Get correlation coefficient between X and Y
func GetCorrelation() float64 {
 res := 0.0
 return res
}