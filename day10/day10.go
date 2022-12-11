package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

func processInstructionPt1(instruction string, x_register *int, x_register_values *[]int) {
	if instruction == "noop" {
		*x_register_values = append(*x_register_values, *x_register)
	} else {
		instruction_split := strings.Split(instruction, " ")
		if instruction_split[0] == "addx" {
			*x_register_values = append(*x_register_values, *x_register)
			value, _ := strconv.Atoi(instruction_split[1])
			*x_register += value
			*x_register_values = append(*x_register_values, *x_register)
		}
	}
}

func appendCrtOutput(x_register *int, crt_output *[]int, crt_position int) {
	if *x_register > crt_position+1 || *x_register < crt_position-1 {
		*crt_output = append(*crt_output, 0)
	} else {
		*crt_output = append(*crt_output, 1)
	}
}

func processInstructionPt2(instruction string, x_register *int, crt_output *[]int) {
	if instruction == "noop" {
		appendCrtOutput(x_register, crt_output, (len(*crt_output))%40)
	} else {
		instruction_split := strings.Split(instruction, " ")
		if instruction_split[0] == "addx" {
			appendCrtOutput(x_register, crt_output, (len(*crt_output))%40)
			appendCrtOutput(x_register, crt_output, (len(*crt_output))%40)
			value, _ := strconv.Atoi(instruction_split[1])
			*x_register += value
		}
	}
}

func calcSignalStrength(cycleCount int, x_register_values []int) int {
	result := cycleCount * x_register_values[cycleCount-2]
	fmt.Printf("Signal Strength at %d: %d\n", cycleCount, result)
	return result
}

func printCrtOutput(crt_output []int) {
	for j := 0; j < 6; j++ {
		for i := 0; i < 40; i++ {
			o := crt_output[i+(j*40)]
			if o == 1 {
				fmt.Print("#")
			} else if o == 0 {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}
}

func processFile(filename string) {
	instructions := loadFile(filename)
	x_register_values := make([]int, 0)
	x_register_pt1 := 1
	x_register_pt2 := 1
	crt_output := make([]int, 0)

	for _, instruction := range instructions {
		processInstructionPt1(instruction, &x_register_pt1, &x_register_values)
		processInstructionPt2(instruction, &x_register_pt2, &crt_output)
	}
	total := 0
	for i := 20; i <= 220; i += 40 {
		total += calcSignalStrength(i, x_register_values)
	}
	fmt.Println("Signal Strength Sum: ", total)

	printCrtOutput(crt_output)
}

func main() {
	processFile("day10-example.txt")
	fmt.Println("-------")
	processFile("day10.txt")
}
