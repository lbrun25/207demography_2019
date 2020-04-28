package demography

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"utils"
)

const (
	notEnoughArgs = "There are not enough arguments.\n"
	doesNotExist = "does not exist.\n"
	fileName = "207demography_data.csv"
	notEnoughValue = "There are not enough value in the file.\n"
	countryCodeDoesNotExist = "The country code does not exist.\n"
	minArg = 1
)

func printErrorWithValue(valueName string, errorMessage string) {
	fmt.Printf("Error: '%s' %s\n", valueName, errorMessage)
}

func printError(errorMessage string) {
	fmt.Printf("Error: %s\n", errorMessage)
}

// CheckHelp arg -h
func CheckHelp() bool {
	argsWithoutProg := os.Args[1:]

	for _, arg := range argsWithoutProg {
		if arg == "-h" {
			return true
		}
	}
	return false
}

func checkOpenFile(fileName string) (bool, *os.File) {
	file, err := os.Open(fileName)
	if err != nil {
		printErrorWithValue(fileName, doesNotExist)
		return false, nil
	}
	return true, file
}

func retrieveValues(countryCode string, file *os.File) bool {
	scan := bufio.NewScanner(file)
	for scan.Scan() {
		values := strings.SplitN(scan.Text(), ";", -1)
		if values[1] == countryCode {
			var intSlice []int
			for i := 2; i < len(values); i++ {
				res := utils.ConvertStringToInt(values[i])
				intSlice = append(intSlice, res)
			}
			if len(intSlice) < 58 {
				printError(notEnoughValue)
				return false
			}
			demography := Demography{
				countryName: values[0],
				countryCode: values[1],
				values:      intSlice,
			}
			Demographies = append(Demographies, demography)
			return true
		}
	}

	if err := scan.Err(); err != nil {
		log.Fatal(err)
		return false
	}
	return false
}

// CheckArgs check user input's args
func CheckArgs() bool {
	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) < minArg {
		fmt.Printf("%s\n", notEnoughArgs)
		return false
	}

	status, file := checkOpenFile(fileName)
	if !status {
		return false
	}
	for _, arg := range argsWithoutProg {
		if !utils.IsCountryCode(arg) {
			fmt.Printf("%s\n", countryCodeDoesNotExist)
			return false
		}
		if !retrieveValues(arg, file) {
			return false
		}
	}
	defer file.Close()
	return true
}