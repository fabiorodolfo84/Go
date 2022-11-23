package main

import "fmt"

func main() {
	var variavel1 string = "variavel 1" // Este primeiro ex. criamos uma variavel com tipo "string"

	variavel2 := "Variavel 2" // este ex: criamos uma variavel sem o tipo mais mais utilizando o simbulo da marmota (:=)

	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var (
		variavel3 string = "Terceiro tipo de variavel" //este é um outro exemplo para declarar variaveis
	)

	variavel5, variavel6, variavel7 := "variavel 5", "variavel 6", "variavel 7" // podemos tambem declarar um mesmo tipo de variavel utilizando a virgula.

	fmt.Println(variavel3)
	fmt.Println(variavel5, variavel6, variavel7)

	const constante1 string = "constante 1" // nesta linha declaramos uma constante que é uma variavel que sera utilizada no programa todo.
	fmt.Println(constante1)

}
