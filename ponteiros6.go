package main

type livro struct {
	titulo string
	autor  string
}

func alterar(l *livro) {

	if l.autor == "Anonimo" {
		l.autor = "Desconhecido"
	}

}
