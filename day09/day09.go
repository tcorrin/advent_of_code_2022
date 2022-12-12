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

type knot struct {
	parent  *position
	current *position
}

type position struct {
	x int
	y int
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

func addCurrentTailPosition(tail *position, tail_positions *[]position) {
	match := false
	for _, p := range *tail_positions {
		if p.x == tail.x && p.y == tail.y {
			match = true
			break
		}
	}
	if !match {
		*tail_positions = append(*tail_positions, *tail)
	}
}

func processCommand(rope *[]*knot, tail_positions *[]position, c command) {
	for i := 0; i < c.amount; i++ {
		moveHead((*rope)[0].current, c.direction)
		for j := 1; j < len(*rope); j++ {
			moveKnot((*rope)[j], c.direction)
		}
		addCurrentTailPosition((*rope)[len(*rope)-1].current, tail_positions)
	}
}

func moveHead(head *position, direction string) {
	if direction == "R" {
		head.x += 1
	} else if direction == "L" {
		head.x -= 1
	} else if direction == "U" {
		head.y += 1
	} else if direction == "D" {
		head.y -= 1
	}
}

func moveKnot(k *knot, direction string) {
	if k.parent.x > k.current.x+1 {
		k.current.x += 1
		if k.parent.y > k.current.y {
			k.current.y += 1
		} else if k.parent.y < k.current.y {
			k.current.y -= 1
		}
	} else if k.parent.x < k.current.x-1 {
		k.current.x -= 1
		if k.parent.y > k.current.y {
			k.current.y += 1
		} else if k.parent.y < k.current.y {
			k.current.y -= 1
		}
	} else if k.parent.y > k.current.y+1 {
		k.current.y += 1
		if k.parent.x > k.current.x {
			k.current.x += 1
		} else if k.parent.x < k.current.x {
			k.current.x -= 1
		}
	} else if k.parent.y < k.current.y-1 {
		k.current.y -= 1
		if k.parent.x > k.current.x {
			k.current.x += 1
		} else if k.parent.x < k.current.x {
			k.current.x -= 1
		}
	}
}

func processFile(filename string, rope_length int) {
	tail_positions := make([]position, 0)
	rope := make([]*knot, 1)
	rope[0] = &knot{parent: nil, current: &position{x: 0, y: 0}}
	for i := 1; i < rope_length; i++ {
		rope = append(rope, &knot{parent: rope[i-1].current, current: &position{x: 0, y: 0}})
	}
	tail_positions = append(tail_positions, *rope[0].current)
	command_list := loadFile(filename)
	for _, c := range command_list {
		processCommand(&rope, &tail_positions, c)
	}
	fmt.Println("Number of unique tail positions: ", len(tail_positions))
}

func main() {
	processFile("day09-example.txt", 2)
	processFile("day09.txt", 2)
	fmt.Println("-------")
	processFile("day09-example.txt", 10)
	processFile("day09-example02.txt", 10)
	processFile("day09.txt", 10)

}
