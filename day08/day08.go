package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

func processRawData(input []string) [][]int {
	result := make([][]int, 0)
	for _, row := range input {
		r := make([]int, 0)
		for _, column := range row {
			col_int, _ := strconv.Atoi(string(column))
			r = append(r, col_int)
		}
		result = append(result, r)
	}
	return result
}

// func checkTree(x int, y int, tree_map [][]int) bool {
// 	result := true
// 	tree_height := tree_map[y][x]
// 	max_y := len(tree_map) - 1
// 	max_x := len(tree_map[0]) - 1
// 	if (y < max_y && tree_map[y+1][x] > tree_height) ||
// 		(x < max_x && tree_map[y][x+1] > tree_height) ||
// 		(y > 0 && tree_map[y-1][x] > tree_height) ||
// 		(x > 0 && tree_map[y][x-1] > tree_height) {
// 		result = false
// 	}
// 	return result
// }

func checkTree(x int, y int, tree_map [][]int) (bool, int) {
	x_pos_blocked := false
	x_neg_blocked := false
	y_pos_blocked := false
	y_neg_blocked := false
	tree_height := tree_map[y][x]
	max_y := len(tree_map) - 1
	max_x := len(tree_map[0]) - 1
	tree_count := 0
	scenic_score := 1
	//fmt.Printf("x: %d y: %d tree_height: %d max_y: %d max_x: %d \n", x, y, tree_height, max_x, max_y)
	for x_pos := x + 1; x_pos <= max_x; x_pos++ {
		tree_count++
		if tree_map[y][x_pos] >= tree_height {
			x_pos_blocked = true
			break
		}
	}
	scenic_score = tree_count * scenic_score
	tree_count = 0
	for x_neg := x - 1; x_neg >= 0; x_neg-- {
		tree_count++
		if tree_map[y][x_neg] >= tree_height {
			x_neg_blocked = true
			break
		}
	}
	scenic_score = tree_count * scenic_score
	tree_count = 0
	for y_pos := y + 1; y_pos <= max_y; y_pos++ {
		tree_count++
		if tree_map[y_pos][x] >= tree_height {
			y_pos_blocked = true
			break
		}
	}
	scenic_score = tree_count * scenic_score
	tree_count = 0
	for y_neg := y - 1; y_neg >= 0; y_neg-- {
		tree_count++
		if tree_map[y_neg][x] >= tree_height {
			y_neg_blocked = true
			break
		}
	}
	scenic_score = tree_count * scenic_score
	vis := !y_neg_blocked || !y_pos_blocked || !x_neg_blocked || !x_pos_blocked
	//fmt.Println("Visible: ", vis)
	return vis, scenic_score
}

func checkTrees(tree_map [][]int) (int, int) {
	visible_trees := 0
	max_y := len(tree_map) - 1
	max_x := len(tree_map[0]) - 1
	best_view_score := 0
	for y := 0; y <= max_y; y++ {
		for x := 0; x <= max_x; x++ {
			visibility, current_view_score := checkTree(x, y, tree_map)

			if current_view_score > best_view_score {
				best_view_score = current_view_score
			}

			if visibility {
				visible_trees++
			}
		}
	}
	return visible_trees, best_view_score
}

func processFile(filename string) {
	raw_data := loadFile(filename)
	tree_map := processRawData(raw_data)
	visible_trees, best_view_score := checkTrees(tree_map)
	fmt.Println("Visible Trees: ", visible_trees)
	fmt.Println("Highest Scenic Score: ", best_view_score)

}

func main() {
	processFile("day08-example.txt")
	fmt.Println("-------")
	processFile("day08.txt")
}
