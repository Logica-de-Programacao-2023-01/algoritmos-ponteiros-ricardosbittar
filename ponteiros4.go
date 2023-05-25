//Escreva uma função em Go que receba um ponteiro para uma variável inteira e atualize o valor da variável
//para a soma dos valores dos seus dois últimos dígitos
//(por exemplo, se o valor inicial da variável for 1234, o novo valor será 3+4=7).

package main

import "fmt"

func updateValue(p *int) {
	ultimoDigito := *p % 10
	*p /= 10
	penultimoDigito := *p % 10
	*p = ultimoDigito + penultimoDigito

}

func main() {
	x := 154
	updateValue(&x)
	fmt.Println(x)
}
