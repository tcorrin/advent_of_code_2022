package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sort"
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

func comparePackets(lhs, rhs []any, pp_index int) *bool {
	var result *bool = nil

	if len(lhs) == 0 && len(rhs) > 0 {
		return newTrue()
	}

	for i, v := range lhs {
		if i > len(rhs)-1 {
			result = newFalse()
			break
		} else {
			lhs_type := fmt.Sprintf("%T", v)
			rhs_type := fmt.Sprintf("%T", rhs[i])
			if lhs_type == "float64" && rhs_type == "float64" {
				lhs := v.(float64)
				rhs := rhs[i].(float64)
				if lhs < rhs {
					result = newTrue()
					break
				} else if lhs > rhs {
					result = newFalse()
					break
				}
			} else if lhs_type == "[]interface {}" && rhs_type == "[]interface {}" {
				result = comparePackets(v.([]any), rhs[i].([]any), pp_index)
				if result != nil {
					break
				}
			} else if lhs_type == "float64" && rhs_type == "[]interface {}" {
				convert := make([]any, 0)
				convert = append(convert, v)
				result = comparePackets(convert, rhs[i].([]any), pp_index)
				if result != nil {
					break
				}
			} else if lhs_type == "[]interface {}" && rhs_type == "float64" {
				convert := make([]any, 0)
				convert = append(convert, rhs[i])
				result = comparePackets(v.([]any), convert, pp_index)
				if result != nil {
					break
				}
			}
		}

		if i+1 == len(lhs) && i+1 < len(rhs) {
			result = newTrue()
			break
		}
	}
	return result
}

func parsePacketPair(pp packetPair, pp_index int) *bool {
	lhs := parsePacket(pp.lhs)
	rhs := parsePacket(pp.rhs)
	return comparePackets(lhs, rhs, pp_index)
}

func parsePacketPairs(pp_list []packetPair) []*bool {
	result := make([]*bool, 0)
	for i, pp := range pp_list {
		result = append(result, parsePacketPair(pp, i))
		//fmt.Println("-------")
	}
	return result
}

func processFile(filename string) {
	pp_list := loadFile(filename)
	correct_order_list := parsePacketPairs(pp_list)
	indices_total := 0
	for i, correct_order_check := range correct_order_list {
		if correct_order_check != nil && *correct_order_check {
			indices_total += (i + 1)
		}
	}
	fmt.Println("Sum of correct order indices: ", indices_total)

	pp_list = append(pp_list, packetPair{lhs: "[[2]]", rhs: "[[6]]"})
	packet_list := make([][]any, 0)
	for _, pp := range pp_list {
		packet_list = append(packet_list, parsePacket(pp.lhs))
		packet_list = append(packet_list, parsePacket(pp.rhs))
	}

	sort.Slice(packet_list, func(i, j int) bool {
		return *comparePackets(packet_list[i], packet_list[j], 0)
	})

	result := 1

	for i, p := range packet_list {
		//fmt.Printf("%02d - %v\n", i+1, p)
		p_string := fmt.Sprintf("%v", p)
		if p_string == "[[6]]" || p_string == "[[2]]" {
			result *= i + 1
		}
	}

	fmt.Println("Distress signal decoder key: ", result)
}

func main() {
	processFile("day13-example.txt")
	processFile("day13.txt")
}
