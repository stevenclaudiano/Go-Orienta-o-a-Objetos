package main

import (
	"banco/contas"
	"fmt"
)

func main() {
	contaDoBruno := contas.ContaPoupanca{}
	contaDoBruno.Depositar(100)
	contaDoBruno.Sacar(5000)
	fmt.Println(contaDoBruno.ObterSaldo())
}
