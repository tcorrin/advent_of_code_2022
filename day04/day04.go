package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type clearanceSection struct {
	upper_bound int
	lower_bound int
}

type elfPair struct {
	elf_a clearanceSection
	elf_b clearanceSection
}

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

func parseElfPairs(input []string) []elfPair {
	result := make([]elfPair, 0)
	for _, raw_ep := range input {
		raw_ep_split := strings.Split(raw_ep, ",")
		result = append(result, elfPair{elf_a: processClearanceSection(raw_ep_split[0]), elf_b: processClearanceSection(raw_ep_split[1])})
	}
	return result
}

func processClearanceSection(input string) clearanceSection {
	input_split := strings.Split(input, "-")
	lower_bound, _ := strconv.Atoi(input_split[0])
	upper_bound, _ := strconv.Atoi(input_split[1])
	return clearanceSection{lower_bound: lower_bound, upper_bound: upper_bound}
}

func checkElfPairForFullyContained(input elfPair) int {
	result := 0
	if input.elf_a.upper_bound <= input.elf_b.upper_bound &&
		input.elf_a.lower_bound >= input.elf_b.lower_bound {
		result = 1
	} else if input.elf_b.upper_bound <= input.elf_a.upper_bound &&
		input.elf_b.lower_bound >= input.elf_a.lower_bound {
		result = 1
	}
	return result
}

func checkElfPairForOverlap(input elfPair) int {
	result := 1
	if input.elf_a.lower_bound > input.elf_b.upper_bound {
		result = 0
	} else if input.elf_b.lower_bound > input.elf_a.upper_bound {
		result = 0
	}
	return result
}

func processFile(filename string) {
	raw_data := loadFile(filename)
	elfPair_list := parseElfPairs(raw_data)
	fullyContainedTotal := 0
	overlapTotal := 0
	for _, ep := range elfPair_list {
		fullyContainedTotal += checkElfPairForFullyContained(ep)
		overlapTotal += checkElfPairForOverlap(ep)
	}
	fmt.Println("Total fully contained matches: ", fullyContainedTotal)
	fmt.Println("Total overlap matches: ", overlapTotal)
}

func main() {
	processFile("day04-example.txt")
	processFile("day04.txt")
}
