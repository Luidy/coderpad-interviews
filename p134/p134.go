package main

import (
	"errors"
	"fmt"
)

type SparseArray interface {
	init(arr []int, size int)
	set(i, val int) error
	get(i int) (int, error)
}

type ImplSparseArray struct {
	l    map[int]int
	size int
}

func (isa *ImplSparseArray) init(arr []int, size int) {
	isa.l = make(map[int]int)
	for i, a := range arr {
		if a != 0 {
			isa.l[i] = a
		}
	}
	isa.size = size
}

func (isa *ImplSparseArray) set(i, val int) error {
	if i >= 0 && i < isa.size {
		isa.l[i] = val
		return nil
	}
	return errors.New("invalid index")
}

func (isa *ImplSparseArray) get(i int) (int, error) {
	if i >= 0 && i < isa.size {
		return isa.l[i], nil
	}
	return 0, errors.New("invalid index")
}

func NewImplSparseArray() SparseArray {
	return &ImplSparseArray{}
}

func main() {
	isa := NewImplSparseArray()
	isa.init([]int{0, 0, 0, 1}, 6)

	for i := -1; i < 10; i++ {
		v, err := isa.get(i)
		fmt.Printf("result of calling get(%d) [value: %d, err: %v]\n", i, v, err)
	}

	for i := -1; i < 10; i++ {
		err := isa.set(i, i)
		fmt.Printf("result of calling set(%d) [err: %v]\n", i, err)
	}

	for i := -1; i < 10; i++ {
		v, err := isa.get(i)
		fmt.Printf("result of calling get(%d) [value: %d, err: %v]\n", i, v, err)
	}
}
