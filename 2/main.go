package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// filename := "sample.txt"
	filename := "input.txt"
	part1(filename)
	// part2(filename)
}

func part1(filename string) {
	fmt.Println("Part 1:")
	data := load_file(filename)
	var reports [][]int
	reports = [][]int{}
	for _, line := range strings.Split(data, "\n") {
		var report []int
		report = []int{}
		for _, level := range strings.Split(line, " ") {
			level_int, _ := strconv.Atoi(level)
			report = append(report, level_int)
		}
		reports = append(reports, report)

	}

	// fmt.Println(reports)
	var safety []int
	safety = []int{}
	for i, report := range reports {
		previous_level := -1
		direction := 0
		// fmt.Println(report)
		for _, level := range report {
			// fmt.Println(previous_level)
			// fmt.Println(reflect.TypeOf(level))

			if previous_level == -1 {
				previous_level = level

			} else {
				// fmt.Println(previous_level, level, direction)
				if previous_level == level {
					safety = append(safety, 0)
					// fmt.Println("unsafe: equal")
					break
				}
				if previous_level > level {
					// 3 levels
					if direction == 1 {
						safety = append(safety, 0)
						// fmt.Println("unsafe: + to -")
						break
					}
					if (previous_level - level) > 3 {
						safety = append(safety, 0)
						// fmt.Println("unsafe: p > l : 3")
						break
					}
					direction = -1
				} else if previous_level < level {
					// 3 levels
					if direction == -1 {
						safety = append(safety, 0)
						// fmt.Println("unsafe: - to +")
						break
					}
					if (level - previous_level) > 3 {
						safety = append(safety, 0)
						// fmt.Println("unsafe: p < l : 3")
						break
					}
					direction = 1
				}
			}

			previous_level = level
		}
		if len(safety)-1 < i {
			// fmt.Println("safe")
			safety = append(safety, 1)
		}
	}
	// fmt.Println(safety)
	answer := 0
	for _, v := range safety {
		answer += v
	}
	fmt.Println("Answer:", answer)
}

func part2(filename string) {
	fmt.Println("Part 2:")
	data := load_file(filename)
	fmt.Println(data)
}

func load_file(filename string) string {
	data, _ := os.ReadFile(filename)
	return string(data)
}

func save_output(filename string, data string) {
	os.WriteFile(filename, []byte(data), 0644)
}
