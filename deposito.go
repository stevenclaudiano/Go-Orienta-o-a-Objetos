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

	fmt.Println(contaDaSilva.saldo)
	status, valor := contaDaSilva.Depositar(2000)
	fmt.Println(status, valor)

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

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Deposito realizado com sucesso", c.saldo
	} else {
		return "Valor do deposito invalido", c.saldo
	}
}
