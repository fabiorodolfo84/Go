package main

import "fmt"

func main() {
	var numero int64 = 1000000000000
	fmt.Println(numero)

	var numero2 uint32 = 10000
	fmt.Println(numero2)

	// alias
	//INT32 = RUNE
	var numero3 rune = 123456
	fmt.Println(numero3)

	// BYTE = UINTB
	var numero4 byte = 123
	fmt.Println(numero4)

	var numeroReal1 float32 = 12345
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 1230000000.45
	fmt.Println(numeroReal2)

}
