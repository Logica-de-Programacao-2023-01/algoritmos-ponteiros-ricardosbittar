//Implemente uma função que receba um ponteiro para uma struct "Livro" com campos "Título", "Autor" e "Preço"
//e que adicione um desconto de 10% no preço do livro

package main

type Livro struct {
	Título string
	Autor  string
	preco  float64
}

func desconto(p *Livro) {
	var x *float64
	var desconto float64

	desconto = p.preco * 0.90
	*x = desconto
}
