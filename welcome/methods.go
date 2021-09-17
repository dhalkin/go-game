package main

//@todo i believe that accept celsius from user and calculate kelvins as output will be nice

import "fmt"

type celsius float64 // Базовый тип float64
type kelvin float64
type fahrenheit float64

const degrees = 273.15

func main() {

	var temperature celsius = 20
	var temperatureK kelvin = 300
	var temperatureF fahrenheit = 100

	fmt.Println(temperature - degrees)

	fmt.Println(kelvinToCelsiusCauseRedeclared(degrees))
	fmt.Println(celsiusToKelvin(degrees))

	// !!! %)
	fmt.Println(kelvin.celsius(temperatureK))
	fmt.Println(temperatureK.celsius())
	fmt.Println(temperatureF.celsius())
}

// kelvinToCelsius converts °K to °C
func kelvinToCelsiusCauseRedeclared(k kelvin) celsius {
	return celsius(k - degrees) // Необходима конвертация типа
}

func celsiusToKelvin(c celsius) kelvin {
	return kelvin(c + degrees)
}

// olala
func (k kelvin) celsius() celsius { // метод celsius для типа kelvin
	return celsius(k - 273.15)
}

func (f fahrenheit) celsius() celsius { // метод celsius для типа kelvin
	return celsius((f - 32.0) * 5.0 / 9.0)
}
