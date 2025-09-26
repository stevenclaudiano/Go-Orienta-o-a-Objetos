package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	var titular string = "João"
	var numeroAgencia int = 589
	var numeroConta int = 123456
	var saldo float64 = 125.50

	fmt.Println(titular, numeroAgencia, numeroConta, saldo)

	var titular2 string = "Steven"
	var numeroAgencia2 int = 555
	var numeroConta2 int = 122222
	var saldo2 float64 = 1300.50

	fmt.Println(titular2, numeroAgencia2, numeroConta2, saldo2)
}
