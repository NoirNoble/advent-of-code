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

type cube_game_round struct {
	red   int
	green int
	blue  int
}

type cube_game struct {
	gameId int
	rounds []cube_game_round
}

const year string = "2023"
const day string = "day02"
const part string = "part01"

var numOfCubesPerColor = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

var colors = []string{
	"red",
	"green",
	"blue",
}

func isNumeric(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
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
	path := year + "/" + day + "/" + part + "/" + input_name

	lines, err := readLines(path)
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	return lines
}

func getGameId(line string) int {
	gameId := strings.Split(line, ":")[0]
	gameId = strings.TrimSpace(gameId)
	gameId = strings.Split(gameId, " ")[1]
	gameIdInt, _ := strconv.Atoi(gameId)

	return gameIdInt
}

func getRounds(line string) []string {
	roundsText := strings.Split(line, ":")[1]
	roundsText = strings.TrimSpace(roundsText)
	rounds := strings.Split(roundsText, ";")

	// Trim whitespace
	for i := 0; i < len(rounds); i++ {
		rounds[i] = strings.TrimSpace(rounds[i])
	}

	return rounds
}

func getRoundDetails(round string) cube_game_round {
	roundCubes := strings.Split(round, ",")
	var cRound cube_game_round
	for i := 0; i < len(roundCubes); i++ {
		roundCubeStr := strings.TrimSpace(roundCubes[i])
		roundCubeData := strings.Split(roundCubeStr, " ")
		numOfCubes, _ := strconv.Atoi(roundCubeData[0])
		colorOfCube := roundCubeData[1]
		switch colorOfCube {
		case "red":
			cRound.red = numOfCubes
		case "green":
			cRound.green = numOfCubes
		case "blue":
			cRound.blue = numOfCubes
		}
	}

	return cRound
}

func buildGameStrut(line string) cube_game {
	var cGame cube_game
	cGame.gameId = getGameId(line)

	roundsArr := getRounds(line)

	var rounds []cube_game_round

	for i := 0; i < len(roundsArr); i++ {
		rDetails := getRoundDetails(roundsArr[i])
		rounds = append(rounds, rDetails)
	}

	cGame.rounds = rounds

	return cGame
}

func isGamePossible(cGame cube_game) bool {
	rounds := cGame.rounds
	for i := 0; i < len(rounds); i++ {
		if rounds[i].red > numOfCubesPerColor["red"] {
			return false
		} else if rounds[i].green > numOfCubesPerColor["green"] {
			return false
		} else if rounds[i].blue > numOfCubesPerColor["blue"] {
			return false
		}
	}

	return true
}

func main() {
	lines := readInputFileLines()

	var sumOfIds int

	for i := 0; i < len(lines); i++ {
		cGame := buildGameStrut(lines[i])
		isGamePossible := isGamePossible(cGame)
		if isGamePossible {
			sumOfIds += cGame.gameId
		}
	}
	fmt.Println(sumOfIds)
}
