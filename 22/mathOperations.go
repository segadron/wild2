package main

import (
	"fmt"
	"math/big"
)

func main() {
	a, _ := big.NewInt(0).SetString("12345678901234567890", 10)
	b, _ := big.NewInt(0).SetString("98765432109876543210", 10)

	// Умножение
	product := new(big.Int).Mul(a, b)
	fmt.Println("Результат умножения:", product)

	// Деление
	quotient := new(big.Rat).SetFrac(a, b)
	fmt.Println("Результат деления:", quotient)

	// Сложение
	sum := new(big.Int).Add(a, b)
	fmt.Println("Результат сложения:", sum)

	// Вычитание
	difference := new(big.Int).Sub(a, b)
	fmt.Println("Результат вычитания:", difference)
}
