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

// FitValues - struct which holds values for the fit
type FitValues struct {
	a float64
	b float64
	rootMeanSquareDeviation float64
	population float64
}

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

func getBMultiplier(b float64) float64 {
	multiplier := 1.0

	if b < 0 {
		multiplier = -1
	}
	return multiplier
}

// ComputeFit1 - get fit1
func ComputeFit1() FitValues {
	a := getA(fit1)
	b := MeanValues - a * MeanYear
	population := a * targetYear + b

	return FitValues{
		a:                       a / math.Pow(10, 6),
		b:                       (b * getBMultiplier(b)) / math.Pow(10, 6),
		rootMeanSquareDeviation: getRootMeanStandardDeviation(fit1, a, b) / math.Pow(10, 6),
		population:              population / math.Pow(10, 6),
	}
}

func ComputeFit2() FitValues {
	a := getA(fit2)
	b := MeanYear - a * MeanValues
	population := (targetYear - b) / a

	return FitValues{
		a:                       a * math.Pow(10, 6),
		b:                       b * getBMultiplier(b),
		rootMeanSquareDeviation: getRootMeanStandardDeviation(fit2, a, b) / math.Pow(10, 6),
		population:              population / math.Pow(10, 6),
	}
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

// GetCorrelation - Get correlation coefficient between X and Y
func GetCorrelation() float64 {
	res := (Fit.one.rootMeanSquareDeviation * math.Pow(10, 6)) /
		(Fit.two.rootMeanSquareDeviation * math.Pow(10, 6))
	return res
}