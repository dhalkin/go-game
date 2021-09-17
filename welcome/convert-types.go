package main

import (
	"fmt"
	"strconv"
)

func main() {

	var pi rune = 960
	var alpha rune = 940
	var omega rune = 969
	var bang byte = 33
	fmt.Print(string(pi), string(alpha), string(omega), string(bang), "\n")

	fmt.Println("===")

	countdown := 10
	str := "Launch in T minus " + strconv.Itoa(countdown) + " seconds."
	fmt.Println(str) // Выводит: Launch in T minus 10 seconds.

	str2 := fmt.Sprintf("Launch in T minus %v seconds.", countdown)
	fmt.Println(str2)

}
