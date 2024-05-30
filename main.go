package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// FahrenheitToCelsius converts a temperature from Fahrenheit to Celsius.
func FahrenheitToCelsius(fahrenheit float64) float64 {
	return (fahrenheit - 32) * 5 / 9
}

// CelsiusToFahrenheit converts a temperature from Celsius to Fahrenheit.
func CelsiusToFahrenheit(celsius float64) float64 {
	return (celsius * 9 / 5) + 32
}
