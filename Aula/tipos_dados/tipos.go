package main

import "fmt"

// string
// bool
// int -> int8, int16, int64 | uint8, uint16, uint64
// float32 | float64
// Bytes | Runes

func main() {
	var nome = "Pedro"
	var isAdmin = true
	idade := 21
	pi := 3.14
	var testeByte byte = 'c'

	fmt.Printf("String: %s \n", nome)
	fmt.Printf("Bool: %t \n", isAdmin)
	fmt.Printf("Idade: %d \n", idade)
	fmt.Printf("Float: %.2f \n", pi)
	fmt.Println("Byte:", testeByte)
}
