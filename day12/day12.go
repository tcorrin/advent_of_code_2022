package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type position struct {
	x int
	y int
}

type graphEntry struct {
	pos    position
	weight int
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

func buildGraphEntryList(hm heightMap, x int, y int) []graphEntry {
	result := make([]graphEntry, 0)
	if x-1 >= 0 && hm.grid[y][x-1] <= hm.grid[y][x]+1 {
		result = append(result, graphEntry{pos: position{x: x - 1, y: y}, weight: hm.grid[y][x] + 1 - hm.grid[y][x-1]})
	}
	if x+1 < len(hm.grid[0]) && hm.grid[y][x+1] <= hm.grid[y][x]+1 {
		result = append(result, graphEntry{pos: position{x: x + 1, y: y}, weight: hm.grid[y][x] + 1 - hm.grid[y][x+1]})
	}
	if y-1 >= 0 && hm.grid[y-1][x] <= hm.grid[y][x]+1 {
		result = append(result, graphEntry{pos: position{x: x, y: y - 1}, weight: hm.grid[y][x] + 1 - hm.grid[y-1][x]})
	}
	if y+1 < len(hm.grid) && hm.grid[y+1][x] <= hm.grid[y][x]+1 {
		result = append(result, graphEntry{pos: position{x: x, y: y + 1}, weight: hm.grid[y][x] + 1 - hm.grid[y+1][x]})
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].weight < result[j].weight
	})
	return result
}

func buildGraph(hm heightMap) map[position][]graphEntry {
	result := make(map[position][]graphEntry)
	for y := 0; y < len(hm.grid); y++ {
		for x := 0; x < len(hm.grid[0]); x++ {
			k := position{x: x, y: y}
			v := buildGraphEntryList(hm, x, y)
			result[k] = v
		}
	}
	return result
}

func pathContainsPosition(path []position, pos position) bool {
	result := false
	for _, p := range path {
		if p.x == pos.x && p.y == pos.y {
			result = true
			break
		}
	}
	return result
}

func findShortestPath(pg map[position][]graphEntry, start position, end position, path []position) []position {
	_, exists := pg[start]
	if !exists {
		return path
	}
	path = append(path, start)
	//fmt.Println(path)
	//fmt.Printf("Start x: %d Start y %d - End x: %d End y: %d\n", start.x, start.y, end.x, end.y)
	if start.x == end.x && start.y == end.y {
		return path
	}
	shortest := make([]position, 0)
	for _, ge := range pg[start] {
		if !pathContainsPosition(path, ge.pos) {
			newPath := findShortestPath(pg, ge.pos, end, path)
			if len(newPath) > 0 {
				if len(shortest) == 0 || (len(newPath) < len(shortest)) {
					shortest = newPath
				}
			}
		}
	}
	return shortest
}

func processFile(filename string) {
	raw_data := loadFile(filename)
	height_map := parseRawData(raw_data)
	//fmt.Println(height_map)
	pathGraph := buildGraph(height_map)
	//fmt.Println(pathGraph)
	path := make([]position, 0)
	shortestPath := findShortestPath(pathGraph, height_map.start, height_map.end, path)
	fmt.Println(shortestPath)
	fmt.Println("The minimum path length is: ", len(shortestPath)-1)
}

func main() {
	processFile("day12-example.txt")
	fmt.Println("-------")
	processFile("day12.txt")
}
