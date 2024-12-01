package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	sample()
	input1()
}

func sample() {
	fmt.Println("Sample:")
	data := load_file("sample1.txt")
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
		fmt.Println(id1[i], id2[i], distance)
	}
	fmt.Println("Answer:", total)
}

func input1() {
	fmt.Println("Input1:")
	data := load_file("input1.txt")
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

func load_file(filename string) string {
	data, _ := os.ReadFile(filename)
	return string(data)
}

func save_output(filename string, data string) {
	os.WriteFile(filename, []byte(data), 0644)
}
