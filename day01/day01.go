package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func loadFile(filename string) []int {
	file, err := os.Open(filename)
	result := make([]int, 0)

	if err == nil {
		sum := 0
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			if scanner.Text() == "" {
				result = append(result, sum)
				sum = 0
			} else {
				calories, _ := strconv.Atoi(scanner.Text())
				sum += calories
			}
		}
		result = append(result, sum)
	} else {
		fmt.Println(err)
		panic("Error opening file")
	}
	return result
}

func calcMax(input []int) int {
	result := 0
	for _, calorieCount := range input {
		if calorieCount > result {
			result = calorieCount
		}
	}
	return result
}

func processFile(filename string) {
	max_calorie_list := loadFile(filename)
	max_calorie_count := calcMax(max_calorie_list)
	fmt.Println("Max calorie count: ", max_calorie_count)
}

func removeFromList(item int, input []int) []int {
	result := make([]int, 0)
	for _, value := range input {
		if value != item {
			result = append(result, value)
		}
	}
	return result

}

func processFilePt2(filename string) {
	max_calorie_list := loadFile(filename)
	number_one := calcMax(max_calorie_list)
	max_calorie_list = removeFromList(number_one, max_calorie_list)
	number_two := calcMax(max_calorie_list)
	max_calorie_list = removeFromList(number_two, max_calorie_list)
	number_three := calcMax(max_calorie_list)

	fmt.Println("Sum of top three: ", number_one+number_two+number_three)
}

func main() {
	processFile("day01-example.txt")
	processFile("day01.txt")
	processFilePt2("day01-example.txt")
	processFilePt2("day01.txt")
}
