package main

import "os"
import "io"

func main() {

	myString := ""
	myArguments := os.Args

	if len(myArguments) == 1 {
		myString = "Give me one argument"
	} else {
		myString = myArguments[1]
	}

	io.WriteString(os.Stdout, myString)
	io.WriteString(os.Stdout, "\n")
}
