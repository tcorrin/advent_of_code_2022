package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type command struct {
	direction string
	amount    int
}

func loadFile(filename string) []command {
	file, err := os.Open(filename)
	result := make([]command, 0)

	if err == nil {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			raw_command_split := strings.Split(scanner.Text(), " ")
			parsed_amount, _ := strconv.Atoi(raw_command_split[1])
			result = append(result, command{direction: raw_command_split[0], amount: parsed_amount})
		}
	} else {
		fmt.Println(err)
		panic("Error opening file")
	}
	return result
}

func processFile(filename string) {
	
	 := loadFile(filename)
}

func main() {
	processFile("day09-example.txt")
	fmt.Println("-------")
	processFile("day09.txt")
}
