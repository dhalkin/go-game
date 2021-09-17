package main

import "bufio"
import "fmt"
import "os"

// get input
func main() {
	var f *os.File
	f = os.Stdin
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		// waiting forever for user input and print everyone
		fmt.Println(">", scanner.Text())
	}
}
