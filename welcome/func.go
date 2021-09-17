package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {

	num := rand.Intn(9)
	future := time.Unix(12622780800, 0)
	countdown, err := strconv.Atoi("10")
	if err != nil {
		fmt.Println("some went wrong")
	}

	fmt.Println(num)
	fmt.Println(future)
	fmt.Println(countdown)

	celsius := kelvinToCelsius(290.6)
	fmt.Println(celsius)
}

// kelvinToCelsius конвертирует °K в °C
func kelvinToCelsius(k float64) float64 { // Объявляет функцию, что принимает параметр и возвращает результат
	k -= 273.15
	return k
}
