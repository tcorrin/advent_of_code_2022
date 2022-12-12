package main

import (
	"bufio"
	"fmt"
	"os"
)

type position struct {
	x int
	y int
}

type heightMap struct {
	start position
	end   position
	grid  [][]int
}

func mapLetter(item rune) int {
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
	return m[item]
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

func parseRawData(input []string) heightMap {
	result := heightMap{}
	for y, line := range input {
		result.grid = append(result.grid, make([]int, 0))
		for x, letter := range line {
			if letter == 'S' {
				result.start = position{x: x, y: y}
				result.grid[y] = append(result.grid[y], mapLetter('a'))
			} else if letter == 'E' {
				result.end = position{x: x, y: y}
				result.grid[y] = append(result.grid[y], mapLetter('z'))
			} else {
				result.grid[y] = append(result.grid[y], mapLetter(letter))
			}
		}
	}
	return result
}

func processFile(filename string) {
	raw_data := loadFile(filename)
	height_map := parseRawData(raw_data)
	fmt.Println(height_map)
	// paths := find_paths(height_map)
	// min_path := len(height_map.grid) * len(height_map.grid[0])
	// for _, path := range paths {
	// 	if len(path) < min_path {
	// 		min_path = len(path)
	// 	}
	// }
	// fmt.Println("The minimum path length is: ", min_path)
}

func main() {
	processFile("day12-example.txt")
	fmt.Println("-------")
	//processFile("day12.txt")
}
