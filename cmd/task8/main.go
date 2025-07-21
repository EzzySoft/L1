package main

import (
	"flag"
	"fmt"
)

const bitWidth = 64

// Устанавливает i-й бит числа num в значение val (0 или 1)
func setBit(num int64, bit, val uint) int64 {
	if bit < 1 || bit > 64 {
		panic("bit должен быть в диапазоне 1-64")
	}
	pos := bit - 1 // bit=1 -> pos=0 (правый)
	mask := int64(1) << pos
	if val == 1 {
		return num | mask
	}
	return num &^ mask
}

// Печатает число в двоичном виде с лидирующими нулями
func printBits(num int64, width uint) {
	for i := int(width) - 1; i >= 0; i-- {
		fmt.Print((num >> i) & 1)
	}
}

func main() {
	var num int64
	var bit, val uint
	flag.Int64Var(&num, "num", 5, "исходное число")
	flag.UintVar(&bit, "bit", 1, "номер бита (1 — самый правый, 64 — самый левый)")
	flag.UintVar(&val, "val", 0, "значение (0 или 1)")
	flag.Parse()

	fmt.Printf("Было:  ")
	printBits(num, bitWidth)
	fmt.Printf(" (%d)\n", num)

	result := setBit(num, bit, val)

	fmt.Printf("Стало: ")
	printBits(result, bitWidth)
	fmt.Printf(" (%d)\n", result)
}
