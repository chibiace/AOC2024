package main

import "os"

func main() {

}

func load_file(filename string) string {
	data, _ := os.ReadFile(filename)
	return string(data)
}

func save_output(filename string, data string) {
	os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
	os.WriteFile(filename, []byte(data), 0644)
}
