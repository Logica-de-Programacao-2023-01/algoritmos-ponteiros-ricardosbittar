package main

import "fmt"

//Escreva uma função que receba um ponteiro para um inteiro e um valor inteiro n.
//A função deve atualizar o valor do inteiro para a soma dos n primeiros números naturais.
func updateValue(p *int, x int) {
	soma := 0

	for i := 0; i <= x; i++ {
		soma += i

	}
	*p = soma

}
func main() {
	y := 21
	var pt *int = &y
	updateValue(pt, y)
	fmt.Println(y)

}
