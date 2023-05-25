// Escreva uma função em Go que receba um ponteiro para uma string e atualize o valor da string para seu reverso.
package main

import "fmt"

func reverso(p *string) {

	runes := []rune(*p)
	slice := len(runes)
	for i, q := 0, slice-1; i < q; i, q = i+1, q-1 {

		runes[i], runes[q] = runes[q], runes[i]
	}
	*p = string(runes)

}
func main() {
	texto := "hello world"
	fmt.Println("antes", texto)
	reverso(&texto)
	fmt.Println("depois", texto)
}
