package list

import "errors"

type ArrayList struct {
	list []int
	size int
}

func (l *ArrayList) Init(size int) error {
	if size > 0 {
		l.list = make([]int, size)
		return nil
	} else {
		return errors.New("Não pode ser criado um ArrayList com tamanho <= 0")
	}
}

func (l *ArrayList) Size() int {
	return l.size
}

func (l *ArrayList) AddEnd(value int) {
	if l.size < len(l.list) {
		l.list[l.size] = value
		l.size++
	}
}
