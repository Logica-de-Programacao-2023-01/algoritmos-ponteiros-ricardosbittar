//Escreva uma função que receba um ponteiro para um inteiro e verifique se esse inteiro é par ou ímpar.
//A função deve atualizar o valor do inteiro para 0 se ele for par ou para 1 se for ímpar

package main

func updatePont(pont *int) bool {
	var x int
	var p *int = &x

	if *p%2 == 0 {
		return true
	} else {

		return false
	}

}
