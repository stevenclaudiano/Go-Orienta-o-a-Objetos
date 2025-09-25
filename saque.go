package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	contaDaSilva := ContaCorrente{}
	contaDaSilva.titular = "Silva"
	contaDaSilva.saldo = 500

	fmt.Println("Saldo anterior em conta:", contaDaSilva.saldo)
	fmt.Println(contaDaSilva.verificasaque(-100))
	fmt.Println("Saldo atual em conta:", contaDaSilva.saldo)

}

func (c *ContaCorrente) verificasaque(valorSaque float64) string {
	podesacar := valorSaque > 0 && valorSaque <= c.saldo
	if podesacar {
		c.saldo -= valorSaque
		return "Saque realizado com sucesso"
	} else {
		return "Saldo insuficiente"
	}

}
