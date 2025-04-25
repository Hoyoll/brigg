package lib

import (
	"fmt"
	"slices"
)

type Barrel[T any] struct {
	Items []T
	freed []int
}

func (i *Barrel[T]) Add(item T) int {
	if len(i.freed) != 0 {
		index := i.freed[0]
		i.Items[index] = item
		i.freed = slices.Clone(i.freed[1:])
		return index
	}
	i.Items = append(i.Items, item)
	return len(i.Items) - 1
}

func (i *Barrel[T]) Get(index int) (*T, error) {
	length := len(i.Items)
	if length < index || length == index {
		return nil, fmt.Errorf("out of bound")
	}
	return &i.Items[index], nil
}

func (i *Barrel[T]) Free(index int) {
	i.freed = append(i.freed, index)
}

func (i *Barrel[T]) Replace(index int, item T) error {
	length := len(i.Items)
	if index > length || length == index {
		return fmt.Errorf("out of bound")
	}
	i.Items[index] = item
	return nil
}
