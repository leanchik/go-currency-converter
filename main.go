package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введите сумму:")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	sum, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Println("Ошибка: нужно ввести число")
		return
	}
	fmt.Println("Сумма:", sum)

	fmt.Print("Выберите направление (toRUB / fromRUB):")
	directionInput, _ := reader.ReadString('\n')
	direction := strings.TrimSpace(directionInput)

	fmt.Print("Выберите валюту(usd, eur, kzt):")
	currencyInput, _ := reader.ReadString('\n')
	currency := strings.TrimSpace(currencyInput)

	const (
		USD = 92.0
		EUR = 100.00
		KZT = 0.20
	)

	var result float64

	if direction == "toRUB" {
		switch currency {
		case "usd":
			result = sum * USD
			fmt.Printf("%.2f USD = %.2f RUB\n", sum, result)
		case "eur":
			result = sum * EUR
			fmt.Printf("%.2f EUR = %.2f RUB\n", sum, result)
		case "kzt":
			result = sum * KZT
			fmt.Printf("%.2f KZT = %.2f RUB\n", sum, result)
		default:
			fmt.Println("Неизвестная валюта")
		}
	} else if direction == "fromRUB" {
		switch currency {
		case "usd":
			result = sum / USD
			fmt.Printf("%.2f RUB = %.2f USD\n", sum, result)
		case "eur":
			result = sum / EUR
			fmt.Printf("%.2f RUB = %.2f EUR\n", sum, result)
		case "kzt":
			result = sum / KZT
			fmt.Printf("%.2f RUB = %.2f KZT\n", sum, result)
		default:
			fmt.Println("Неизвестная валюта")
		}
	} else {
		fmt.Println("Неизвестное направление")
	}
}
