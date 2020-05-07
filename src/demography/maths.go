package demography

import (
	"fmt"
	"math"
)

// MeanYear - mean of years
var MeanYear float64

// MeanYear - mean of values
var MeanValues float64

type fitValues struct {
	alpha float64
	beta float64
	rmsd float64
	pred float64
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

// ComputeFit1 - get fit1
func ComputeFit1() {
	dividend := 0.0
	divisor := 0.0
	y := 0.0
	idx := 0

	for _, population := range Populations {
		dividend += (float64(population.value) - MeanValues) * (float64(population.year) - MeanYear)
		divisor += math.Pow(float64(population.year) - MeanYear, 2)
	}
	beta := dividend / divisor
	alpha := MeanValues - beta * MeanYear

	dividend = 0
	for _, year := range Years {
		for _, demography := range Demographies {
			y += float64(demography.values[idx])
		}
		dividend += math.Pow((beta * float64(year) + alpha) - y, 2)
		y = 0
		idx++
	}
	rsmd := math.Sqrt(dividend / float64(len(Years)))
	pred := beta * 2050 + alpha
	Fit = Fits{one: fitValues{
		alpha: alpha,
		beta:  beta,
		rmsd:  rsmd,
		pred:  pred,
	}}
	fmt.Printf("alpha = %f\n", Fit.one.alpha)
	fmt.Printf("beta = %f\n", Fit.one.beta)
	fmt.Printf("rmsd = %f\n", Fit.one.rmsd)
	fmt.Printf("pred = %f\n", Fit.one.pred)
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