package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type monkey struct {
	object_list          []uint64
	operation            string
	test                 int
	positive_destination int
	negative_destiantion int
	inspection_count     int
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

func parseMonkeys(input []string) []monkey {
	result := make([]monkey, 0)
	var m *monkey = nil
	for _, line := range input {
		if strings.Contains(line, "Monkey") {
			if m != nil {
				result = append(result, *m)
			}
			m = &monkey{}
		} else if strings.Contains(line, "Starting items:") {
			re := regexp.MustCompile("^  Starting items: (.*)$")
			starting_list_split := strings.Split(re.FindStringSubmatch(line)[1], ", ")
			m.object_list = make([]uint64, 0)
			for _, o := range starting_list_split {
				v, _ := strconv.Atoi(o)
				m.object_list = append(m.object_list, uint64(v))
			}
		} else if strings.Contains(line, "Operation:") {
			re := regexp.MustCompile("^  Operation: new = old (.*)$")
			m.operation = re.FindStringSubmatch(line)[1]
		} else if strings.Contains(line, "Test:") {
			re := regexp.MustCompile("^  Test: divisible by (.*)$")
			match := re.FindStringSubmatch(line)[1]
			m.test, _ = strconv.Atoi(match)
		} else if strings.Contains(line, "If true:") {
			re := regexp.MustCompile("^    If true: throw to monkey (.*)$")
			match := re.FindStringSubmatch(line)[1]
			m.positive_destination, _ = strconv.Atoi(match)
		} else if strings.Contains(line, "If false:") {
			re := regexp.MustCompile("^    If false: throw to monkey (.*)$")
			match := re.FindStringSubmatch(line)[1]
			m.negative_destiantion, _ = strconv.Atoi(match)
		}
	}
	result = append(result, *m)
	return result
}

func processRound(monkeys *[]monkey, pt2 bool) {
	for i, _ := range *monkeys {
		processMonkey(monkeys, i, pt2)
	}
}

func runOperation(input uint64, operation string) uint64 {
	result := uint64(0)
	operation_split := strings.Split(operation, " ")
	var operator_value = input
	if operation_split[1] != "old" {
		parse_result, _ := strconv.Atoi(operation_split[1])
		operator_value = uint64(parse_result)
	}
	if operation_split[0] == "*" {
		result = operator_value * input
	} else if operation_split[0] == "+" {
		result = operator_value + input
	}
	return result
}

func processMonkey(monkeys *[]monkey, index int, pt2 bool) {
	for _, o := range (*monkeys)[index].object_list {
		//fmt.Printf("Monkey %d inspects object - %d\n", index, o)
		new_value := runOperation(o, (*monkeys)[index].operation)
		//fmt.Printf("  Post operation %s - %d\n", (*monkeys)[index].operation, new_value)
		if !pt2 {
			new_value = new_value / 3
			//fmt.Printf("  Post boredom divison by 3 - %d\n", new_value)
		}
		if new_value%uint64((*monkeys)[index].test) == 0 {
			//fmt.Printf("  Item divisible by %d\n", (*monkeys)[index].test)
			//fmt.Printf("  Item %d passed to monkey %d\n", new_value, (*monkeys)[index].positive_destination)
			(*monkeys)[(*monkeys)[index].positive_destination].object_list = append((*monkeys)[(*monkeys)[index].positive_destination].object_list, new_value)
		} else {
			//fmt.Printf("  Item not divisible by %d\n", (*monkeys)[index].test)
			//fmt.Printf("  Item %d passed to monkey %d\n", new_value, (*monkeys)[index].negative_destiantion)
			(*monkeys)[(*monkeys)[index].negative_destiantion].object_list = append((*monkeys)[(*monkeys)[index].negative_destiantion].object_list, new_value)
		}
		(*monkeys)[index].inspection_count++
	}
	(*monkeys)[index].object_list = make([]uint64, 0)
}

func printMonkeyItems(monkeys *[]monkey) {
	for i, m := range *monkeys {
		fmt.Println("Monkey ", i)
		fmt.Println("  Objects: ", m.object_list)
	}
}

func processFile(filename string, pt2 bool) {
	count := 20
	if pt2 {
		count = 1000
	}
	raw_data := loadFile(filename)
	monkeys := parseMonkeys(raw_data)
	for rc := 0; rc < count; rc++ {
		processRound(&monkeys, pt2)
		//printMonkeyItems(&monkeys)
		//fmt.Println("********")
	}
	printMonkeyItems(&monkeys)
	first := 0
	second := 0
	for i, m := range monkeys {
		if m.inspection_count > first {
			if first > second {
				second = first
			}
			first = m.inspection_count
		} else if m.inspection_count > second {
			second = m.inspection_count
		}
		fmt.Printf("Monkey %d inspected %d items\n", i, m.inspection_count)
	}
	fmt.Println(first)
	fmt.Println(second)
	fmt.Println("Monkey business level: ", first*second)
}

func main() {
	//processFile("day11-example.txt", false)
	processFile("day11-example.txt", true)
	fmt.Println("-------")
	//processFile("day11.txt", false)
	//processFile("day11.txt", true)
}
