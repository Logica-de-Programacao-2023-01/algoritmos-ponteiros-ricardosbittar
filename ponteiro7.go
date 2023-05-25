//Escreva uma função em Go que receba um ponteiro para um struct Conta com campos saldo e titular
//e adicione um valor ao saldo da conta

package main

type conta struct {
	titular string
	saldo   float64
}

func alterarSaldo(p *conta) {
	var x *float64
	var valor float64

	valor = p.saldo
	*x = valor

}
