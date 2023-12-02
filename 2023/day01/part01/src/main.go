package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const demo = false

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

func calculateCoordinates(line string) int {
	lineArray := strings.Split(line, "")
	length := len(lineArray)

	var first_num_found bool = false
	var last_num_found bool = false
	var first_num string
	var last_num string

	// Search forward
	for i := 0; i < length; i++ {
		char := string(lineArray[i])
		if first_num_found == false {
			if isNumeric(char) {
				first_num = char
				first_num_found = true
			}
		}
	}

	// Search backwards
	for i := length - 1; i >= 0; i-- {
		char := string(lineArray[i])
		if last_num_found == false {
			if isNumeric(char) {
				last_num = char
				last_num_found = true
			}
		}
	}

	number, _ := strconv.Atoi(first_num + last_num)

	fmt.Sprint(first_num)
	fmt.Sprint(last_num)
	return number
}

func main() {
	// Load file
	var input_name string

	if demo == true {
		input_name = "demo-input.txt"
	} else {
		input_name = "input.txt"
	}
	path := "2023/day01/part01/" + input_name

	lines, err := readLines(path)
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	var calibrationValueSum int
	// Read lines into Array
	for i := 0; i < len(lines); i++ {
		// chars := strings.Split(lines[i], "")
		number := calculateCoordinates(lines[i])
		calibrationValueSum += number
		fmt.Printf("%d: %d\n", i+1, number)
	}
	fmt.Printf("Calibration Value: %d\n", calibrationValueSum)
}
