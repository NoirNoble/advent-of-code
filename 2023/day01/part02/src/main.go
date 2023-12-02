package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

const demo = false

type coordinate_digits struct {
	digit int
	index int
}

var stringNumbers = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func isNumeric(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func findStrNumbers(s string) []coordinate_digits {
	foundNumbers := []coordinate_digits{}
	var first_coordinate_digit coordinate_digits
	var last_coordinate_digit coordinate_digits
	first_coordinate_digit.index = -1
	last_coordinate_digit.index = -1

	// find first occurance
	for strDigit, intDigit := range stringNumbers {
		digitIndex := strings.Index(s, strDigit)
		if digitIndex >= 0 && (first_coordinate_digit.index == -1 || first_coordinate_digit.index > digitIndex) {
			first_coordinate_digit.digit = intDigit
			first_coordinate_digit.index = digitIndex
		}
	}

	// find last occurance
	for strDigit, intDigit := range stringNumbers {
		digitIndex := strings.LastIndex(s, strDigit)
		if digitIndex >= 0 && last_coordinate_digit.index < digitIndex {
			last_coordinate_digit.digit = intDigit
			last_coordinate_digit.index = digitIndex
		}
	}

	if first_coordinate_digit.index != -1 {
		foundNumbers = append(foundNumbers, first_coordinate_digit)
	}

	if last_coordinate_digit.index != -1 {
		foundNumbers = append(foundNumbers, last_coordinate_digit)
	}

	sort.Slice(foundNumbers, func(i, j int) bool {
		return foundNumbers[i].index < foundNumbers[j].index
	})

	fmt.Println(foundNumbers)
	return foundNumbers
}

func findIntNumbers(s string) []coordinate_digits {
	foundNumbers := []coordinate_digits{}
	lineArray := strings.Split(s, "")
	length := len(lineArray)

	var first_num_found bool = false
	var last_num_found bool = false

	// Search forward
	for i := 0; i < length; i++ {
		char := string(lineArray[i])
		if first_num_found == false {
			if isNumeric(char) {
				first_num_found = true
				var coordinate_digit coordinate_digits
				coordinate_digit.index = i
				coordinate_digit.digit, _ = strconv.Atoi(char)
				foundNumbers = append(foundNumbers, coordinate_digit)
			}
		}
	}

	// Search backwards
	for i := length - 1; i >= 0; i-- {
		char := string(lineArray[i])
		if last_num_found == false {
			if isNumeric(char) {
				last_num_found = true
				var coordinate_digit coordinate_digits
				coordinate_digit.index = i
				coordinate_digit.digit, _ = strconv.Atoi(char)
				foundNumbers = append(foundNumbers, coordinate_digit)
			}
		}
	}

	return foundNumbers
}

func calculateCoordinates(intNumbers []coordinate_digits, strNumbers []coordinate_digits) int {
	mixedNumbers := append(intNumbers, strNumbers...)
	sort.Slice(mixedNumbers, func(i, j int) bool {
		return mixedNumbers[i].index < mixedNumbers[j].index
	})

	stringDigit := fmt.Sprintf("%d%d", mixedNumbers[0].digit, mixedNumbers[len(mixedNumbers)-1].digit)
	intDigit, _ := strconv.Atoi(stringDigit)
	return intDigit
}

func main() {
	// Load file
	var input_name string

	if demo == true {
		input_name = "demo-input.txt"
	} else {
		input_name = "input.txt"
	}
	path := "2023/day01/part02/" + input_name

	lines, err := readLines(path)
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	var calibrationValueSum int
	// Read lines into Array
	for i := 0; i < len(lines); i++ {
		strNumbers := findStrNumbers(lines[i])
		intNumbers := findIntNumbers(lines[i])
		lineDigit := calculateCoordinates(strNumbers, intNumbers)

		calibrationValueSum += lineDigit

		fmt.Printf("\"%s\" = %v\n", lines[i], lineDigit)

	}

	fmt.Printf("Calibration Value: %d\n", calibrationValueSum)
}
