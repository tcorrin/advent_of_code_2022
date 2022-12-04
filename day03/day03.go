package main

import (
	"bufio"
	"fmt"
	"os"
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

func mapItem(item rune) int {
	m := make(map[rune]int)
	m['a'] = 1
	m['b'] = 2
	m['c'] = 3
	m['d'] = 4
	m['e'] = 5
	m['f'] = 6
	m['g'] = 7
	m['h'] = 8
	m['i'] = 9
	m['j'] = 10
	m['k'] = 11
	m['l'] = 12
	m['m'] = 13
	m['n'] = 14
	m['o'] = 15
	m['p'] = 16
	m['q'] = 17
	m['r'] = 18
	m['s'] = 19
	m['t'] = 20
	m['u'] = 21
	m['v'] = 22
	m['w'] = 23
	m['x'] = 24
	m['y'] = 25
	m['z'] = 26
	m['A'] = 27
	m['B'] = 28
	m['C'] = 29
	m['D'] = 30
	m['E'] = 31
	m['F'] = 32
	m['G'] = 33
	m['H'] = 34
	m['I'] = 35
	m['J'] = 36
	m['K'] = 37
	m['L'] = 38
	m['M'] = 39
	m['N'] = 40
	m['O'] = 41
	m['P'] = 42
	m['Q'] = 43
	m['R'] = 44
	m['S'] = 45
	m['T'] = 46
	m['U'] = 47
	m['V'] = 48
	m['W'] = 49
	m['X'] = 50
	m['Y'] = 51
	m['Z'] = 52
	return m[item]
}

func processRucksack(contents_list string) int {
	sep := len(contents_list) / 2
	comp_one := contents_list[:sep]
	comp_two := contents_list[sep:]
	for _, item_one := range comp_one {
		for _, item_two := range comp_two {
			if item_one == item_two {
				return mapItem(item_one)
			}
		}
	}
	panic("No matching item")
}

func processGroup(rucksack_list []string) int {
	ruck_one := rucksack_list[0]
	ruck_two := rucksack_list[1]
	ruck_three := rucksack_list[2]

	for _, item_one := range ruck_one {
		for _, item_two := range ruck_two {
			if item_one == item_two {
				for _, item_three := range ruck_three {
					if item_one == item_three {
						return mapItem(item_one)
					}
				}
			}
		}
	}
	panic("No matching item")
}

func processFile(filename string) {
	rucksack_list := loadFile(filename)
	total := 0
	for _, ruck := range rucksack_list {
		total += processRucksack(ruck)
	}

	fmt.Println("Total score: ", total)
}

func processFilePt2(filename string) {
	rucksack_list := loadFile(filename)
	total := 0
	for x := 0; x < len(rucksack_list); x += 3 {
		total += processGroup(rucksack_list[x : x+3])
	}

	fmt.Println("Total score pt2: ", total)
}

func main() {
	processFile("day03-example.txt")
	processFile("day03.txt")
	processFilePt2("day03-example.txt")
	processFilePt2("day03.txt")
}
