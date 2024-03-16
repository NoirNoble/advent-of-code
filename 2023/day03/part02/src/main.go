package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const demo = true

type part struct {
	value      int
	startIndex int
	endIndex   int
	row        int
}

var searchPattern = [][]int{
	{-1, -1}, // top left
	{-1, 0},  // top centre
	{-1, 1},  // top right
	{0, -1},  // mid left
	{0, 1},   // mid right
	{1, -1},  // bottom left
	{1, 0},   // bottom centre
	{1, 1},   // bottom right
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

func readInputFileLines() []string {
	var input_name string
	if demo == true {
		input_name = "demo-input.txt"
	} else {
		input_name = "input.txt"
	}
	// path := year + "/" + day + "/" + part + "/" + input_name
	path := "../" + input_name

	lines, err := readLines(path)
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	return lines
}

func isNumeric(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func createMatrix(lines []string) [][]string {
	matrix := make([][]string, len(lines))
	var row []string

	for i := 0; i < len(lines); i++ {
		row = strings.Split(lines[i], "")
		matrix[i] = row
	}

	return matrix
}

func getNumbersAndIndexes(matrix [][]string) []part {
	var parts []part

	for i := range matrix {
		row := matrix[i]
		var currentNum []string
		var rowParts []part
		var searchingNumber bool = false
		var startIndex int

		for j := range row {
			var part part
			if isNumeric(row[j]) {
				if !searchingNumber {
					startIndex = j
				}
				searchingNumber = true
				currentNum = append(currentNum, row[j])
				if j == len(row)-1 {
					part.startIndex = startIndex
					part.endIndex = j - 1
					numberString := strings.Join(currentNum, "")
					part.value, _ = strconv.Atoi(numberString)
					part.row = i
					rowParts = append(rowParts, part)
					currentNum = nil
					searchingNumber = false
				}
			} else if searchingNumber {
				part.startIndex = startIndex
				part.endIndex = j - 1
				numberString := strings.Join(currentNum, "")
				part.value, _ = strconv.Atoi(numberString)
				part.row = i
				rowParts = append(rowParts, part)
				currentNum = nil
				searchingNumber = false
			}

		}

		parts = append(parts, rowParts...)

	}

	return parts
}

func checkIfPartNumber(part part, matrix [][]string) bool {

	row := part.row
	for column := part.startIndex; column <= part.endIndex; column++ {
		for index, element := range searchPattern {
			if row == 0 && (index == 0 || index == 1 || index == 2) {
				// if row = 0 {!searchPattern[0, 1, 2]}
				continue
			} else if row == len(matrix)-1 && (index == 5 || index == 6 || index == 7) {
				// if row = len(matrix) {!searchPattern[5, 6, 7]}
				continue
			} else if column == 0 && (index == 0 || index == 3 || index == 5) {
				// if column = 0 {!searchPattern[0, 3, 5]}
				continue
			} else if column == len(matrix[row])-1 && (index == 2 || index == 4 || index == 7) {
				// if column = len(row) {!searchPattern[2, 4, 7]}
				continue
			}
			adjacentChar := matrix[row+element[0]][column+element[1]]
			if !isNumeric(adjacentChar) && adjacentChar != "." {
				return true
			}
		}
	}

	return false

}

func getGearIdentifier(matrix [][]string) [][]int {
	var gears [][]int
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == "*" {
				gears = append(gears, []int{i, j})
			}
		}
	}

	return gears
}

func rangeGen(start int, end int) []int {
	array := make([]int, end-start+1)
	for i := range array {
		array[i] = i + start
	}

	return array
}

func isCoordInArray(coord []int, array [][]int) bool {
	for _, element := range array {
		if element[0] == coord[0] && element[1] == coord[1] {
			return true
		}
	}

	return false
}

func getGearPower(gears [][]int, parts []part) int {
	var totalGearPower int
	for _, gear := range gears {
		gearRow := gear[0]
		gearCol := gear[1]
		var gearSearchIndexes [][]int
		var partNumbers []part

		// Find gear search indexes
		for _, element := range searchPattern {
			row := gearRow + element[0]
			col := gearCol + element[1]
			gearSearchIndexes = append(gearSearchIndexes, []int{row, col})
		}

		// Find if parts are in those search indexes
		for _, part := range parts {
			// find all coordinates for part
			partRange := rangeGen(part.startIndex, part.endIndex)
			var partCoordRange [][]int
			for _, e := range partRange {
				partCoordRange = append(partCoordRange, []int{part.row, e})
			}

			// Check if any of part's coordinates are in range
			for _, partCoord := range partCoordRange {
				if isCoordInArray(partCoord, gearSearchIndexes) {
					if len(partNumbers) >= 2 {
						partNumbers = nil
					} else {
						partNumbers = append(partNumbers, part)
					}

				}
			}

		}

		// Calculate power
		if len(partNumbers) == 2 {
			gearPower := partNumbers[0].value * partNumbers[1].value
			totalGearPower += gearPower
		}

	}

	return totalGearPower
}

func main() {
	lines := readInputFileLines()
	matrix := createMatrix(lines)
	// // fmt.Println(matrix[0][1]) // [row][column]
	// var total int

	// for _, element := range numbers {
	// 	if checkIfPartNumber(element, matrix) {
	// 		total += element.value
	// 	}
	// }

	// f, err := os.Create("code-file.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// defer f.Close()

	// row := 0
	// for _, element := range numbers {
	// 	if element.row > row {
	// 		f.WriteString("\n")
	// 		row = element.row
	// 	}
	// 	f.WriteString(fmt.Sprintf("%d,", element.value))
	// }

	// fmt.Println(numbers)
	// fmt.Println(getGearIdentifier(matrix))

	parts := getNumbersAndIndexes(matrix)
	gears := getGearIdentifier(matrix)
	total := getGearPower(gears, parts)
	// getGearPower(gears, parts)
	// var partCoordRange [][]int
	// partRange := rangeGen(numbers[3].startIndex, numbers[3].endIndex)
	// for _, e := range partRange {
	// 	partCoordRange = append(partCoordRange, []int{numbers[3].row, e})
	// }
	// fmt.Println(partCoordRange)

	fmt.Println(total)

	// fmt.Println(checkIfPartNumber(numbers[1], matrix))
}

// Save numbers
// Save index of first and last character
// Check adjacent characters
// if symbol, save if not discard
