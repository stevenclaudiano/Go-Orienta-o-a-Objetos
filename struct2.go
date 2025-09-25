package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {

	contajoao := ContaCorrente{"João", 123, 456789, 1000.50}
	contajoao2 := ContaCorrente{"João", 123, 4567890, 1000.50}

	contamaria := ContaCorrente{titular: "Maria", numeroAgencia: 421, numeroConta: 987654, saldo: 2500.75}
	contamaria2 := ContaCorrente{titular: "Maria", numeroAgencia: 321, numeroConta: 987654, saldo: 2500.75}
	fmt.Println(contajoao == contajoao2)
	fmt.Println(contamaria == contamaria2)

	var contaDaCris *ContaCorrente
	contaDaCris = new(ContaCorrente)
	contaDaCris.titular = "Cris"
	contaDaCris.saldo = 300

	//fmt.Println(*contaDaCris)

	var contaDaCris2 *ContaCorrente
	contaDaCris = new(ContaCorrente)
	contaDaCris.titular = "Cris"
	contaDaCris.saldo = 300

	fmt.Println(&contaDaCris)
	fmt.Println(&contaDaCris2)

	fmt.Println(*contaDaCris == *contaDaCris2)
}
