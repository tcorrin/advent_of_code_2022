package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type craneOperation struct {
	origin      string
	destination string
	quantity    int
}

func loadFile(filename string) ([]string, []string) {
	file, err := os.Open(filename)
	crate_stacks := make([]string, 0)
	crane_operations := make([]string, 0)
	operations := false

	if err == nil {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			if scanner.Text() == "" {
				operations = true
			} else if !operations {
				crate_stacks = append(crate_stacks, scanner.Text())
			} else {
				crane_operations = append(crane_operations, scanner.Text())
			}
		}
	} else {
		fmt.Println(err)
		panic("Error opening file")
	}
	return crate_stacks, crane_operations
}

func parseCraneOperations(input []string) []craneOperation {
	result := make([]craneOperation, 0)
	for _, raw_co := range input {
		raw_co_split := strings.Split(raw_co, " ")
		q, _ := strconv.Atoi(raw_co_split[1])
		result = append(result, craneOperation{origin: raw_co_split[3], destination: raw_co_split[5], quantity: q})
	}
	return result
}

func parseCrateStacks(input []string) map[string]string {
	result := make(map[string]string)
	stack_labels := input[len(input)-1]
	stacks := input[:len(input)-1]
	for label_index, label := range stack_labels {
		if label != ' ' {
			result[string(label)] = parseStack(label_index, stacks)
		}
	}
	return result
}

func parseStack(index int, stack []string) string {
	result := ""
	for _, level := range stack {
		if level[index] != ' ' {
			result = result + string(level[index])
		}
	}
	return reverseString(result)
}

func reverseString(input string) (result string) {
	for _, v := range input {
		result = string(v) + result
	}
	return
}

func runOperation(operation craneOperation, stacks map[string]string, part_two bool) map[string]string {
	offset := len(stacks[operation.origin]) - operation.quantity
	pickup := stacks[operation.origin][offset:]
	if !part_two {
		pickup = reverseString(pickup)
	}
	stacks[operation.destination] = stacks[operation.destination] + pickup
	stacks[operation.origin] = stacks[operation.origin][:offset]
	return stacks
}

func calcTopRow(stacks map[string]string) string {
	result := ""
	for i := 1; i <= len(stacks); i++ {
		result += string(stacks[strconv.Itoa(i)][len(stacks[strconv.Itoa(i)])-1])

	}
	return result
}

func processFile(filename string) {
	raw_crate_stacks, raw_crane_operations := loadFile(filename)
	crate_stacks := parseCrateStacks(raw_crate_stacks)
	crate_stacks_pt2 := parseCrateStacks(raw_crate_stacks)
	crane_operations := parseCraneOperations(raw_crane_operations)

	for _, o := range crane_operations {
		runOperation(o, crate_stacks, false)
	}

	for _, o := range crane_operations {
		runOperation(o, crate_stacks_pt2, true)
	}
	fmt.Println("Top Row: ", calcTopRow(crate_stacks))
	fmt.Println("Top Row Pt2: ", calcTopRow(crate_stacks_pt2))
}

func main() {
	processFile("day05-example.txt")
	processFile("day05.txt")
}
