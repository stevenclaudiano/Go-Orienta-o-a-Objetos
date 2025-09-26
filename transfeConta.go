package main

import (
	"Banco/Contas" // nomeDoModulo/nomeDaPasta
	"fmt"
)

func main() {
	contaDaSilva := Contas.ContaCorrente{titular: "Silva", saldo: 500.0}
	contaDoJoao := Contas.ContaCorrente{titular: "João", saldo: 300.0}

	status := contaDaSilva.tranferir(-200, &contaDoJoao)
	fmt.Println(status)

	fmt.Println(contaDaSilva)
	fmt.Println(contaDoJoao)

}
