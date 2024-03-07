package main

import (
	"eds/list"
	"fmt"
)

func main() {
	// Criando a lista
	size := 10
	arraylist := list.ArrayList{}
	arraylist.Init(size)

	// Adicionando valores
	for i := 0; i < size; i++ {
		arraylist.AddEnd(i)
	}

	// Imprimindo
	fmt.Print(arraylist)

	// Teste AddEnd()
	arraylist.AddEnd(5)
	fmt.Print(arraylist)

	// Teste GetByIndex()
	arraylist.GetByIndex(2)
}
