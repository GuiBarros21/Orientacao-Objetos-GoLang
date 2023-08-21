package Contas

type ContaCorrente struct {
	Titular       string
	NumeroAgencia int
	NumeroConta   int
	Saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) (string, float64) {
	podeSacar := valorDoSaque <= c.Saldo && valorDoSaque > 0
	if podeSacar {
		c.Saldo -= valorDoSaque
		return "Saque realizado com sucesso", c.Saldo

	} else {
		return "Saldo insifuciente", c.Saldo
	}
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) string {
	podeSacar := valorDoDeposito > 0
	if podeSacar {
		c.Saldo += valorDoDeposito
		return "Deposito realizado com sucesso."
	} else {
		return "O valor do deposito menor que zero."
	}
}

func (c *ContaCorrente) Transferir(valorDaTransfencia float64, contaDestino *ContaCorrente) bool {
	if valorDaTransfencia < c.Saldo {
		c.Saldo -= valorDaTransfencia
		contaDestino.Depositar(valorDaTransfencia)
		return true
	} else {
		return false
	}
}
