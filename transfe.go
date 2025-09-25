package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	contaDaSilva := ContaCorrente{titular: "Silva", saldo: 500.0}
	contaDoJoao := ContaCorrente{titular: "João", saldo: 300.0}

	status := contaDaSilva.tranferir(-200, &contaDoJoao)
	fmt.Println(status)

	fmt.Println(contaDaSilva)
	fmt.Println(contaDoJoao)

}

func (c *ContaCorrente) tranferir(valorDaTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorDaTransferencia < c.saldo && valorDaTransferencia > 0 {
		c.saldo -= valorDaTransferencia
		contaDestino.Depositar(valorDaTransferencia)
		return true
	} else {
		return false
	}
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
