//Implemente uma função que receba um ponteiro para uma slice e seu tamanho
//e preencha-o com os n primeiros números primos.

package main

func primos(num int) bool {

	if num <= 1 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func colocarPrimos(slice *[]int, tamanho int) {
	nprimos := 0
	numeros := 2

	for nprimos < tamanho {
		if primos(numeros) {
			*slice = append(*slice, numeros)
			nprimos++

		}
		numeros++

	}
}
