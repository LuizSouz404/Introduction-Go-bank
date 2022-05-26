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
		conta.saldo -= saque

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
		conta.saldo += deposito

		fmt.Println("O saldo atual da sua conta R$", conta.saldo)
	}
}

func (conta *Account) transferir(destino *Account) {
	var transferencia float64

	fmt.Println("Quanto deseja transferir em R$")
	fmt.Scan(&transferencia)

	valorValido := transferencia > 0 && conta.saldo > transferencia

	if valorValido {
		conta.saldo -= transferencia

		destino.saldo += transferencia

		fmt.Println("O saldo atual da sua conta R$", conta.saldo)
		fmt.Println("Tranferido R$", transferencia)
	}
}

func main() {
	contaDoLuiz := Account{titular: "Luiz", numeroAgencia: 589, numeroConta: 213465, saldo: 3125.55}
	contaDaBrunna := Account{titular: "Brunna", numeroAgencia: 589, numeroConta: 213465, saldo: 3125.55}

	fmt.Println(contaDoLuiz)

	fmt.Println("=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=")
	fmt.Println("=+=+=+=+   Conta bancaria   =+=+=+=")
	fmt.Println("=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=")

	for {
		fmt.Println("\nSelecione uma opção")
		fmt.Println("[1] Vizualizar saldo\n[2] Sacar saldo \n[3] Depositar\n[4] Transferir\n[0] Sair")
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
			contaDoLuiz.transferir(&contaDaBrunna)
			fmt.Println(contaDaBrunna.saldo)
		case 0:
			os.Exit(0)
		default:
			os.Exit(-1)
		}
	}
}
