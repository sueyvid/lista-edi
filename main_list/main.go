package main

import (
	"eds/list"
	"fmt"
)

func main() {
	size := 10
	arraylist := list.ArrayList{}
	arraylist.Init(size)
	for i := 0; i < size; i++ {
		arraylist.AddEnd(i)
	}
	fmt.Print(arraylist)
	arraylist.AddEnd(5)
	fmt.Print(arraylist)
}
