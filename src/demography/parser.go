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
	countryCodeInvalidFormat = "Invalid format for the country code.\n"
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

func retrieveFileContent(file *os.File) (bool, *[]string) {
	scan := bufio.NewScanner(file)
	var lines []string

	for scan.Scan() {
		lines = append(lines, scan.Text())
	}
	if err := scan.Err(); err != nil {
		log.Fatal(err)
		return false, nil
	}
	return true, &lines
}

func retrieveValues(countryCode string, lines []string) bool {
	for _, line := range lines {
		values := strings.SplitN(line, ";", -1)
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
	printErrorWithValue(countryCode, countryCodeDoesNotExist)
	return false
}

// CheckArgs check user input's args
func CheckArgs() bool {
	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) < minArg {
		printError(notEnoughArgs)
		return false
	}

	status, file := checkOpenFile(fileName)
	if !status {
		return false
	}
	status, lines := retrieveFileContent(file)
	if !status {
		return false
	}
	for _, arg := range argsWithoutProg {
		if !utils.IsCountryCode(arg) {
			printErrorWithValue(arg, countryCodeInvalidFormat)
			return false
		}
		if !retrieveValues(arg, *lines) {
			return false
		}
	}
	defer file.Close()
	return true
}