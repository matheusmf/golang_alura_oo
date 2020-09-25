package main

import (
	"banco/clientes"
	"banco/contas"
	"fmt"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {
	clienteDenis := clientes.Titular{"Denis", "123.123.123.12", "Desenvolvedor"}
	dadosContaDenis := contas.DadosConta{Titular: clienteDenis}
	contaDoDenis := contas.ContaPoupanca{DadosConta: dadosContaDenis}
	fmt.Println(contaDoDenis)
	contaDoDenis.Depositar(100)
	PagarBoleto(&contaDoDenis, 60)

	fmt.Println(contaDoDenis.ObterSaldo())

	clienteLuisa := clientes.Titular{Nome: "Luisa"}
	dadosContaLuisa := contas.DadosConta{Titular: clienteLuisa, NumeroAgencia: 111}
	contaDaLuisa := contas.ContaCorrente{DadosConta: dadosContaLuisa}
	fmt.Println(contaDaLuisa)
	contaDaLuisa.Depositar(500)
	PagarBoleto(&contaDaLuisa, 1000)

	fmt.Println(contaDaLuisa.ObterSaldo())
}

// func main() {
// 	clienteBruno := clientes.Titular{"Bruno", "123.123.123.12", "Desenvolvedor"}
// 	contaDoBruno := contas.ContaCorrente{clienteBruno, 123, 123456, 100}
// 	fmt.Println(contaDoBruno)
// }

// func main() {
// 	contaDaSilvia := contas.ContaCorrente{}
// 	contaDaSilvia.Titular = "Silvia"
// 	contaDaSilvia.Saldo = 500

// 	fmt.Println(contaDaSilvia.Saldo)
// 	status, valor := contaDaSilvia.Depositar(2000)
// 	fmt.Println(status, valor)

// 	contaDoGustavo := contas.ContaCorrente{Titular: "Gustavo", Saldo: 100}

// 	fmt.Println(contaDaSilvia)
// 	fmt.Println(contaDoGustavo)

// 	statusT1 := contaDaSilvia.Transferir(200, &contaDoGustavo)
// 	fmt.Println(statusT1)
// 	statusT2 := contaDoGustavo.Transferir(200, &contaDaSilvia)
// 	fmt.Println(statusT2)
// 	statusT3 := contaDoGustavo.Transferir(-200, &contaDaSilvia)
// 	fmt.Println(statusT3)
// }
