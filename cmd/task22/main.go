package main

import (
	"fmt"
	"math/big"
)

func main() {
	// Задаём большие числа в виде строк
	aStr := "123456789012345678901234567890"
	bStr := "987654321098765432109876543210"

	// Создаём big.Int из строк
	a := new(big.Int)
	b := new(big.Int)
	a.SetString(aStr, 10) // 10 — система счисления
	b.SetString(bStr, 10)

	// Сложение
	sum := new(big.Int).Add(a, b)
	fmt.Println("Сумма:", sum.String())

	// Вычитание
	diff := new(big.Int).Sub(a, b)
	fmt.Println("Разность:", diff.String())

	// Умножение
	prod := new(big.Int).Mul(a, b)
	fmt.Println("Произведение:", prod.String())

	// Деление (для наглядности поменял местами делимое и делитель)
	quot, mod := new(big.Int), new(big.Int)
	quot.DivMod(b, a, mod)
	fmt.Println("Целая часть:", quot.String())
	fmt.Println("Остаток:", mod.String())
}
