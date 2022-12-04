package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func loadFile(filename string) []string {
	file, err := os.Open(filename)
	result := make([]string, 0)

	if err == nil {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			result = append(result, scanner.Text())
		}
	} else {
		fmt.Println(err)
		panic("Error opening file")
	}
	return result
}

func processRock(player string) int {
	result := 0
	if player == "X" {
		result = 3 + 1
	} else if player == "Y" {
		result = 6 + 2
	} else if player == "Z" {
		result = 0 + 3
	}
	return result
}

func processPaper(player string) int {
	result := 0
	if player == "X" {
		result = 0 + 1
	} else if player == "Y" {
		result = 3 + 2
	} else if player == "Z" {
		result = 6 + 3
	}
	return result
}

func processScissors(player string) int {
	result := 0
	if player == "X" {
		result = 6 + 1
	} else if player == "Y" {
		result = 0 + 2
	} else if player == "Z" {
		result = 3 + 3
	}
	return result
}

func processRockPt2(player string) int {
	result := 0
	if player == "X" {
		result = 0 + 3
	} else if player == "Y" {
		result = 3 + 1
	} else if player == "Z" {
		result = 6 + 2
	}
	return result
}

func processPaperPt2(player string) int {
	result := 0
	if player == "X" {
		result = 0 + 1
	} else if player == "Y" {
		result = 3 + 2
	} else if player == "Z" {
		result = 6 + 3
	}
	return result
}

func processScissorsPt2(player string) int {
	result := 0
	if player == "X" {
		result = 0 + 2
	} else if player == "Y" {
		result = 3 + 3
	} else if player == "Z" {
		result = 6 + 1
	}
	return result
}

func processRound(round string) (int, int) {
	result := 0
	resultPt2 := 0
	round_split := strings.Split(round, " ")
	opp := round_split[0]
	player := round_split[1]
	if opp == "A" {
		result = processRock(player)
		resultPt2 = processRockPt2(player)
	} else if opp == "B" {
		result = processPaper(player)
		resultPt2 = processPaperPt2(player)
	} else if opp == "C" {
		result = processScissors(player)
		resultPt2 = processScissorsPt2(player)
	}
	return result, resultPt2
}

func processFile(filename string) {
	round_list := loadFile(filename)
	total_score := 0
	total_score_pt2 := 0
	for _, round := range round_list {
		result, result_pt2 := processRound(round)
		total_score += result
		total_score_pt2 += result_pt2
	}
	fmt.Println("Total Score: ", total_score)
	fmt.Println("Total Score Pt2: ", total_score_pt2)
}

func main() {
	processFile("day02-example.txt")
	processFile("day02.txt")
}
