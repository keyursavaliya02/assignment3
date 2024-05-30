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
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	for {
		value, unit, err := promptUser()
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		switch unit {
		case "F":
			celsius := FahrenheitToCelsius(value)
			fmt.Printf("%.2f F is %.2f C\n", value, celsius)
		case "C":
			fahrenheit := CelsiusToFahrenheit(value)
			fmt.Printf("%.2f C is %.2f F\n", value, fahrenheit)
		}

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Do you want to convert another temperature? (y/n): ")
		response, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading response:", err)
			continue
		}
		response = strings.TrimSpace(response)
		if strings.ToLower(response) != "y" {
			break
		}
	}

	fmt.Println("Exit!")
}
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

// promptUser gets the user's input for the temperature and the unit.
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

func main() {
	for {
		value, unit, err := promptUser()
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		switch unit {
		case "F":
			celsius := FahrenheitToCelsius(value)
			fmt.Printf("%.2f F is %.2f C\n", value, celsius)
		case "C":
			fahrenheit := CelsiusToFahrenheit(value)
			fmt.Printf("%.2f C is %.2f F\n", value, fahrenheit)
		}

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Do you want to convert another temperature? (y/n): ")
		response, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading response:", err)
			continue
		}
		response = strings.TrimSpace(response)
		if strings.ToLower(response) != "y" {
			break
		}
	}

	fmt.Println("Exit")
}
