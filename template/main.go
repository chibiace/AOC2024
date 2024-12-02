package main

import (
	"fmt"
	"os"
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
	fmt.Println(data)
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
