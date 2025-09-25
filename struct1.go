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
	contamaria := ContaCorrente{titular: "Maria", numeroAgencia: 321, numeroConta: 987654, saldo: 2500.75}
	fmt.Println(contajoao)
	fmt.Println(contamaria)

	var contaDaCris *ContaCorrente
	contaDaCris = new(ContaCorrente)
	contaDaCris.titular = "Cris"
	contaDaCris.saldo = 300

	fmt.Println(contaDaCris)

}

// ctrl + ;
