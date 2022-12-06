package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	example01 string = "mjqjpqmgbljsphdztnvjfqwrcgsmlb"
	example02 string = "bvwbjplbgvbhsrlpgdmjqwftvncz"
	example03 string = "nppdvjthqldpwncqszvftbrmjlhg"
	example04 string = "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"
	example05 string = "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"
)

func loadFile(filename string) string {
	result := ""
	file, err := os.Open(filename)
	if err == nil {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			result = scanner.Text()
		}
	} else {
		fmt.Println(err)
		panic("Error opening file")
	}
	return result
}

func allCharactersUnique(input string) bool {
	result := true
	for _, char := range input {
		test := strings.ReplaceAll(input, string(char), "")
		if len(test) != len(input)-1 {
			result = false
			break
		}
	}
	return result
}

func processDataStream(input string, msize int) {
	for index := range input {
		if index >= msize-1 && allCharactersUnique(input[index-(msize-1):index+1]) {
			fmt.Println("first start of packet marker after character ", index+1)
			break
		}
	}
}

func main() {
	processDataStream(example01, 4)
	processDataStream(example02, 4)
	processDataStream(example03, 4)
	processDataStream(example04, 4)
	processDataStream(example05, 4)
	processDataStream(loadFile("day06.txt"), 4)
	processDataStream(example01, 14)
	processDataStream(example02, 14)
	processDataStream(example03, 14)
	processDataStream(example04, 14)
	processDataStream(example05, 14)
	processDataStream(loadFile("day06.txt"), 14)
}
