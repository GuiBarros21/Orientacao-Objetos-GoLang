package main

import (
	"fmt"
)

func main() {

	contaDoGuilherme := ContaCorrente{
		"Guilherme",
		589,
		12345,
		500.45,
	}

	contaDoGuilherme2 := ContaCorrente{
		"Guilherme",
		589,
		12345,
		123.45,
	}
	contaDoGuilherme.Transferir(100, &contaDoGuilherme2)

	fmt.Println(contaDoGuilherme)
	fmt.Println(contaDoGuilherme2)
	// var contaDaCris *ContaCorrente
	// contaDaCris = new(ContaCorrente)
	// contaDaCris.titular = "Cris"
	// contaDaCris.saldo = 500

}
