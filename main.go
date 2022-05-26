package main

import (
	"fmt"
	"os"
)

type Account struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func (conta *Account) VisualizarSaldo() {
	fmt.Println("O saldo da sua conta R$", conta.saldo)
}

func (conta *Account) SacarSaldo() {
	var saque float64

	fmt.Println("Quanto deseja sacar em R$")
	fmt.Scan(&saque)

	saqueValido := saque < conta.saldo && 0 < saque

	if saqueValido {
		conta.saldo = conta.saldo - saque

		fmt.Println("O saldo atual da sua conta R$", conta.saldo)
	} else {
		fmt.Println("Operação de saque inválida, saque maior que o saldo atual R$", conta.saldo)
	}
}

func (conta *Account) DepositarSaldo() {
	var deposito float64

	fmt.Println("Quanto deseja depositar em R$")
	fmt.Scan(&deposito)

	valorValido := deposito > 0

	if valorValido {
		conta.saldo = conta.saldo + deposito

		fmt.Println("O saldo atual da sua conta R$", conta.saldo)
	}
}

func main() {
	contaDoLuiz := Account{titular: "Luiz", numeroAgencia: 589, numeroConta: 213465, saldo: 3125.55}

	fmt.Println(contaDoLuiz)

	fmt.Println("=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=")
	fmt.Println("=+=+=+=+   Conta bancaria   =+=+=+=")
	fmt.Println("=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=")

	for {
		fmt.Println("\nSelecione uma opção")
		fmt.Println("[1] Vizualizar saldo\n[2] Sacar saldo \n[3] Depositar\n[4] Sair")
		var opt int

		fmt.Scan(&opt)

		switch opt {
		case 1:
			contaDoLuiz.VisualizarSaldo()
		case 2:
			contaDoLuiz.SacarSaldo()
		case 3:
			contaDoLuiz.DepositarSaldo()
		case 4:
			os.Exit(0)
		default:
			os.Exit(-1)
		}
	}
}
