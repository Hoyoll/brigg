package lib

import (
	"fmt"
)

type Barrel[T any] struct {
	Items []T
	freed map[int]struct{}
}

func (i *Barrel[T]) Add(item T) int {
	if len(i.freed) != 0 {
		for k, _ := range i.freed {
			i.Items[k] = item
			delete(i.freed, k)
			return k
		}
	}
	i.Items = append(i.Items, item)
	return len(i.Items) - 1
}

func (i *Barrel[T]) Get(index int) (*T, error) {
	if index >= len(i.Items) || index < 0 {
		var zero T
		return &zero, fmt.Errorf("out of bound")
	}
	return &i.Items[index], nil
}

func (i *Barrel[T]) Free(index int) {
	var zero T
	i.Items[index] = zero
	i.freed[index] = struct{}{}
}

func (i *Barrel[T]) Replace(index int, item T) error {
	length := len(i.Items)
	if index > length || length == index {
		return fmt.Errorf("out of bound")
	}
	i.Items[index] = item
	return nil
}

func (i *Barrel[T]) Iter(f func(*T) bool) {
	for k := range i.Items {
		_, ok := i.freed[k]
		if ok {
			continue
		}
		if !f(&i.Items[k]) {
			break
		}
	}
}
