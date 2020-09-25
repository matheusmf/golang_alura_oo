package contas

import "banco/clientes"

type DadosConta struct {
	Titular                    clientes.Titular
	NumeroAgencia, NumeroConta int
}
