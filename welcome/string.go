package main

import "fmt"

func main() {
	question := "¿Cómo estás?"

	for i, c := range question {
		fmt.Printf("%v %c-%v\n", i, c, c)
	}
	fmt.Println("====")

	for _, c := range question {
		fmt.Printf("%c ", c)
	}
	fmt.Println("\n====")

	// rune byte
	var grade rune = 'A'
	fmt.Printf("%v \n", grade)
	fmt.Printf("%c \n", grade)
	fmt.Println("====")

	message := "shalom"
	c := message[0]
	fmt.Printf("%c\n", c) // Выводит: m

}
