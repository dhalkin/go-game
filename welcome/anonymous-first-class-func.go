package main

import (
	"fmt"
	"math/rand"
	"time"
)

type kelvinRedeclared float64

// declare func
var sensor func() kelvinRedeclared

// also declare type function
//type sensor func() kelvinRedeclared

func fakeSensor() kelvinRedeclared {
	return kelvinRedeclared(rand.Intn(151) + 150)
}

func realSensor() kelvinRedeclared {
	return 0 // TO-DO: внедрить реальный сенсор
}

func measureTemperature(samples int, sensor func() kelvinRedeclared) { // measureTemperature принимает функцию в качестве второго параметра
	for i := 0; i < samples; i++ {
		k := sensor()
		fmt.Printf("%v° K\n", k)
		time.Sleep(time.Second)
	}
}

func main() {
	sensor := fakeSensor // Присваивает функцию переменной
	fmt.Println(sensor())

	sensor = realSensor
	fmt.Println(sensor())

	// transfer func to another func
	measureTemperature(3, fakeSensor)
}
