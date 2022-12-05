package main

import (
	"bufio"
	"fmt"
	"os"
)

type craneOperation struct {
	origin      int
	destination int
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
		append(result, craneOperation{}
	}
	return result
}

func processFile(filename string) {
	raw_crate_stacks, raw_crane_operations  := loadFile(filename)
	crate_stacks := parseCrateStacks(raw_crate_stacks)
	crane_operations := parseCraneOperations(raw_crane_operations)



}

func main() {
	processFile("day05-example.txt")
	processFile("day05.txt")
}
