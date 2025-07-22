package main

import "fmt"

func printType(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	case chan int, chan string, chan bool, chan interface{}:
		fmt.Println("chan")
	default:
		fmt.Println("unknown type")
	}
}

func main() {
	printType(42)                // int
	printType("hello")           // string
	printType(true)              // bool
	printType(make(chan int))    // chan
	printType(make(chan string)) // chan
	printType(1.23)              // unknown type
}
