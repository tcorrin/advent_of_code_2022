package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	total_file_system_size int = 70000000
	update_space_required  int = 30000000
)

type file struct {
	name string
	size int
}

type directory struct {
	parent_dir  *directory
	name        string
	directories *[]directory
	files       *[]file
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

func findDirectory(input *[]directory, lookup string) *directory {
	for _, d := range *input {
		if d.name == lookup {
			return &d
		}
	}
	panic("Could not find directory " + lookup)
}

func processInstructions(input []string) directory {
	directories := make([]directory, 0)
	files := make([]file, 0)
	base_dir := directory{directories: &directories, files: &files, name: "base_dir", parent_dir: nil}
	current_dir := &base_dir
	for _, i := range input[1:] {
		if i[0:1] == "$" {
			if i[2:4] == "cd" {
				dir := i[5:]
				if dir == ".." {
					current_dir = current_dir.parent_dir
				} else {
					current_dir = findDirectory(current_dir.directories, dir)
				}
			}
		} else {
			if i[0:3] == "dir" {
				directories := make([]directory, 0)
				files := make([]file, 0)
				new_dir := directory{directories: &directories, files: &files, parent_dir: current_dir, name: i[4:]}
				if current_dir.directories == nil {

					current_dir.directories = &directories
				}
				*current_dir.directories = append(*current_dir.directories, new_dir)
			} else {
				fileSplit := strings.Split(i, " ")
				size, _ := strconv.Atoi(fileSplit[0])
				new_file := file{size: size, name: fileSplit[1]}
				*current_dir.files = append(*current_dir.files, new_file)
			}
		}
	}

	return base_dir
}

func findDirectorySize(input directory, currentSize *int) *int {
	for _, f := range *input.files {
		*currentSize += f.size
	}
	for _, d := range *input.directories {
		currentSize = findDirectorySize(d, currentSize)
	}
	return currentSize
}

func findPuzzleOutput(input directory, currentOutput *int) *int {
	for _, d := range *input.directories {
		fds_input := 0
		fds_output := *findDirectorySize(d, &fds_input)
		if fds_output <= 100000 {
			*currentOutput += fds_output
		}
		currentOutput = findPuzzleOutput(d, currentOutput)
	}
	return currentOutput
}

func findDeletionCandidates(input directory, minimum_space_to_free int, deletion_candidates *[]int) *[]int {
	for _, d := range *input.directories {
		fds_input := 0
		fds_output := *findDirectorySize(d, &fds_input)
		if fds_output >= minimum_space_to_free {
			*deletion_candidates = append(*deletion_candidates, fds_output)
		}
		deletion_candidates = findDeletionCandidates(d, minimum_space_to_free, deletion_candidates)
	}
	return deletion_candidates
}

func processFile(filename string) {
	instruction_list := loadFile(filename)
	base_dir := processInstructions(instruction_list)
	pt1_output := 0
	fmt.Println("Pt1 output: ", *findPuzzleOutput(base_dir, &pt1_output))
	total_base_dir_size := 0
	total_base_dir_size = *findDirectorySize(base_dir, &total_base_dir_size)
	available_space := total_file_system_size - total_base_dir_size
	minimum_space_to_free := update_space_required - available_space
	fmt.Println("Total base dir size: ", total_base_dir_size)
	fmt.Println("Space Required: ", update_space_required)
	fmt.Println("Space Available: ", available_space)
	fmt.Println("Minimum Space To Free: ", minimum_space_to_free)
	deletion_candidates := make([]int, 0)
	deletion_candidates = *findDeletionCandidates(base_dir, minimum_space_to_free, &deletion_candidates)
	min_value := total_file_system_size
	for _, dc := range deletion_candidates {
		if dc < min_value {
			min_value = dc
		}
	}
	fmt.Println("Size of smallest folder to delete: ", min_value)

}

func main() {
	processFile("day07-example.txt")
	fmt.Println("-------")
	processFile("day07.txt")
}
