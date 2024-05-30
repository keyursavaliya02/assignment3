func FahrenheitToCelsius(fahrenheit float64) float64 {
	return (fahrenheit - 32) * 5 / 9
}

func CelsiusToFahrenheit(celsius float64) float64 {
	return (celsius * 9 / 5) + 32
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func promptUser() (float64, string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter temperature (e.g., 100 C or 32 F): ")
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, "", fmt.Errorf("error reading input")
	}

	input = strings.TrimSpace(input)
	parts := strings.Fields(input)
	if len(parts) != 2 {
		return 0, "", fmt.Errorf("invalid input format")
	}

	value, err := strconv.ParseFloat(parts[0], 64)
	if err != nil {
		return 0, "", fmt.Errorf("invalid temperature value")
	}

	unit := strings.ToUpper(parts[1])
	if unit != "C" && unit != "F" {
		return 0, "", fmt.Errorf("invalid temperature unit")
	}

	return value, unit, nil
}
