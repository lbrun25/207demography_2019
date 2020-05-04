package demography

var meanYear float64
var meanValues float64

// ComputeYearMean - get mean date (year)
func ComputeYearMean() {
	sum := 0.0

	for i := 0; i < len(Years); i++ {
		sum += float64(Years[i])
	}
	meanYear = sum /  float64(len(Years))
}

// ComputeMeanValues - get mean demography values
func ComputeMeanValues() {
	sum := 0.0

	for _, demography := range Demographies {
		for _, value := range demography.values {
			sum += float64(value)
		}
	}
	meanValues = (sum / float64(len(Years))) / 100
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