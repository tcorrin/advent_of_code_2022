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

func queuePop(q *[]position) position {
	result := (*q)[0]
	(*q)[0] = position{}
	*q = (*q)[1:]
	return result
}

func findShortestPath(pg map[position][]graphEntry, start position, end position) []position {
	path := make([]position, 0)
	exploredNodes := make(map[position]position, 0)
	queue := make([]position, 0)
	exploredNodes[start] = position{}
	queue = append(queue, start)
	success := false
	for len(queue) > 0 {
		current := queuePop(&queue)
		if current.x == end.x && current.y == end.y {
			success = true
			break
		}
		for _, next_step := range pg[current] {
			_, exists := exploredNodes[next_step.pos]
			if !exists {
				queue = append(queue, next_step.pos)
				exploredNodes[next_step.pos] = current
			}
		}

	}
	if success {
		current := end
		for current.x != start.x || current.y != start.y {
			path = append(path, current)
			current = exploredNodes[current]
			//printGrid(hm, path)
		}
	}
	return path
}

// func printGrid(hm heightMap, path []position) [][]string {
// 	cmd := exec.Command("clear")
// 	cmd.Stdout = os.Stdout
// 	cmd.Run()
// 	g := make([][]string, 0)
// 	for y := 0; y < len(hm.grid); y++ {
// 		g = append(g, make([]string, 0))
// 		for x := 0; x < len(hm.grid[0]); x++ {
// 			g[y] = append(g[y], fmt.Sprintf("%02d", hm.grid[y][x]))
// 		}
// 	}
// 	g[hm.start.y][hm.start.x] = "S "
// 	g[hm.end.y][hm.end.x] = " E"

// 	for _, pos := range path {
// 		g[pos.y][pos.x] = "\033[32m##\033[0m"
// 	}

// 	for y := 0; y < len(hm.grid); y++ {
// 		fmt.Println(g[y])
// 	}
// 	time.Sleep(100 * time.Millisecond)
// 	return g
// }

func processFile(filename string) {
	raw_data := loadFile(filename)
	hm := parseRawData(raw_data)
	pathGraph := buildGraph(hm)
	shortestPath := findShortestPath(pathGraph, hm.start, hm.end)
	fmt.Println("Pt1 - The minimum path length is: ", len(shortestPath))
	base_elevation_list := make([]position, 0)
	pt2_path_list := make([][]position, 0)
	for y := 0; y < len(hm.grid); y++ {
		for x := 0; x < len(hm.grid[0]); x++ {
			if hm.grid[y][x] == 1 {
				base_elevation_list = append(base_elevation_list, position{x: x, y: y})
			}
		}
	}
	fmt.Println("Number of paths to check: ", len(base_elevation_list))
	for _, s := range base_elevation_list {
		pt2_path_list = append(pt2_path_list, findShortestPath(pathGraph, s, hm.end))
	}
	min_path_length := 50000
	for _, p := range pt2_path_list {
		if len(p) != 0 && len(p) < min_path_length {
			min_path_length = len(p)
		}
	}
	fmt.Println("Pt2 - The minimum path length is: ", min_path_length)
}

func main() {
	processFile("day12-example.txt")
	processFile("day12.txt")
}
