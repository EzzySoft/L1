package main

import "fmt"

// Обмен значений a и b через XOR без третьей переменной
func main() {
	a, b := 3, 7

	a = a ^ b
	fmt.Printf("a = a ^ b: a = %03b, b = %03b\n", a, b)

	b = a ^ b
	fmt.Printf("b = a ^ b: a = %03b, b = %03b\n", a, b)

	a = a ^ b
	fmt.Printf("a = a ^ b: a = %03b, b = %03b\n", a, b)

	fmt.Printf("Результат: a = %d, b = %d\n", a, b)
}
