package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type packetPair struct {
	lhs string
	rhs string
}

func newTrue() *bool {
	b := true
	return &b
}

func newFalse() *bool {
	b := false
	return &b
}

func loadFile(filename string) []packetPair {
	result := make([]packetPair, 0)
	pp := packetPair{lhs: "", rhs: ""}
	file, err := os.Open(filename)
	if err == nil {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			t := scanner.Text()
			if t == "" {
				result = append(result, pp)
				pp = packetPair{lhs: "", rhs: ""}
			} else if pp.lhs == "" {
				pp.lhs = t
			} else if pp.rhs == "" {
				pp.rhs = t
			}
		}
		result = append(result, pp)
	} else {
		fmt.Println(err)
		panic("Error opening file")
	}
	return result
}

func parsePacket(input string) []any {
	var arr []any
	err := json.Unmarshal([]byte(input), &arr)
	if err != nil {
		panic(err.Error())
	}
	return arr
}

func compareList(lhs, rhs []any, pp_index int) *bool {
	var result *bool = nil

	if pp_index+1 == 50 {
		fmt.Println("~~TROUBLESOME NUMBER 50~~")
		fmt.Println("LHS")
		fmt.Println(lhs[0])
		fmt.Println("RHS")
		fmt.Println(rhs[0])
	}

	if len(lhs) == 0 && len(rhs) > 0 {
		fmt.Println("~~EVIL NUMBER SIX~~")
		fmt.Println("LHS ran out of items - PASS - index ", pp_index+1)
		return newTrue()
	}

	for i, v := range lhs {
		fmt.Printf("i: %d LHS len: %d RHS len: %d\n", i, len(lhs), len(rhs))
		if i > len(rhs)-1 {
			fmt.Println("RHS ran out of items - FAIL - index ", pp_index+1)
			result = newFalse()
			break
		} else if i+1 == len(lhs)-1 && i+1 < len(rhs)-1 {
			fmt.Println("LHS ran out of items - PASS - index ", pp_index+1)
			result = newTrue()
			break
		} else {
			lhs_type := fmt.Sprintf("%T", v)
			rhs_type := fmt.Sprintf("%T", rhs[i])
			if lhs_type == "float64" && rhs_type == "float64" {
				lhs := v.(float64)
				rhs := rhs[i].(float64)
				if lhs < rhs {
					fmt.Println("LHS smaller - PASS - index ", pp_index+1)
					result = newTrue()
					break
				} else if lhs > rhs {
					fmt.Println("RHS smaller - FAIL - index ", pp_index+1)
					result = newFalse()
					break
				}
			} else if lhs_type == "[]interface {}" && rhs_type == "[]interface {}" {
				result = compareList(v.([]any), rhs[i].([]any), pp_index)
				if result != nil {
					break
				}
			} else if lhs_type == "float64" && rhs_type == "[]interface {}" {
				convert := make([]any, 0)
				convert = append(convert, v)
				result = compareList(convert, rhs[i].([]any), pp_index)
				if result != nil {
					break
				}
			} else if lhs_type == "[]interface {}" && rhs_type == "float64" {
				convert := make([]any, 0)
				convert = append(convert, rhs[i])
				result = compareList(v.([]any), convert, pp_index)
				if result != nil {
					break
				}
			}
		}
		if pp_index+1 == 50 {
			fmt.Println("EXIT")
			fmt.Println(lhs)
			fmt.Println(rhs)
		}
	}
	return result
}

func parsePacketPair(pp packetPair, pp_index int) *bool {
	lhs := parsePacket(pp.lhs)
	rhs := parsePacket(pp.rhs)
	return compareList(lhs, rhs, pp_index)
}

func parsePacketPairs(pp_list []packetPair) []*bool {
	result := make([]*bool, 0)
	for i, pp := range pp_list {
		result = append(result, parsePacketPair(pp, i))
		fmt.Println("-------")
	}
	return result
}

func processFile(filename string) {
	pp_list := loadFile(filename)
	//fmt.Println(pp_list)
	correct_order_list := parsePacketPairs(pp_list)
	for i, c := range correct_order_list {
		if c != nil {
			fmt.Printf("%d - %t\n", i+1, *c)
		} else {
			fmt.Printf("%d - nil\n", i+1)
		}
	}
	indices_total := 0
	for i, correct_order_check := range correct_order_list {
		if correct_order_check != nil && *correct_order_check {
			indices_total += (i + 1)
		}
	}
	fmt.Println("Sum of correct order indices: ", indices_total)
}

func main() {
	//processFile("day13-example.txt")
	processFile("day13.txt")
}
