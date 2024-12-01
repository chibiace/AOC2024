package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// filename := "sample.txt"
	filename := "input.txt"
	part1(filename)
	part2(filename)
}

func part1(filename string) {
	fmt.Println("Part 1:")
	data := load_file(filename)
	var id1 []int
	var id2 []int
	for _, line := range strings.Split(data, "\n") {
		split := strings.Split(line, "   ")
		v1, _ := strconv.Atoi(split[0])
		v2, _ := strconv.Atoi(split[1])
		id1 = append(id1, v1)
		id2 = append(id2, v2)
	}
	sort.Ints(id1)
	sort.Ints(id2)
	var total int
	total = 0
	for i, _ := range id1 {
		distance := id1[i] - id2[i]
		if distance < 0 {
			distance = -distance
		}
		total += distance
		// fmt.Println(id1[i], id2[i], distance)
	}
	fmt.Println("Answer:", total)
}

func part2(filename string) {
	fmt.Println("Part 2:")
	data := load_file(filename)
	var id1 []int
	var id2 []int
	for _, line := range strings.Split(data, "\n") {
		split := strings.Split(line, "   ")
		// fmt.Println(split)
		v1, _ := strconv.Atoi(split[0])
		v2, _ := strconv.Atoi(split[1])
		id1 = append(id1, v1)
		id2 = append(id2, v2)
	}
	total := 0
	for _, v1 := range id1 {
		times := 0
		for _, v2 := range id2 {
			if v1 == v2 {
				times += 1
			}
		}
		score := v1 * times
		total += score
		// fmt.Println(v1, times, score)
	}
	fmt.Println("Answer:", total)
}

func load_file(filename string) string {
	data, _ := os.ReadFile(filename)
	return string(data)
}

func save_output(filename string, data string) {
	os.WriteFile(filename, []byte(data), 0644)
}
