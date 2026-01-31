package main

import "fmt"

type DATable[T any] struct {
	tables []*T
}

func (t *DATable[T]) Search(k int) (T, bool) {
	if k < 0 || k > len(t.tables) {
		// return zero value
		var zero T
		return zero, false
	} else {
		return *t.tables[k], true
	}
}

func (t *DATable[T]) Insert(v T, k int) {
	if k < 0 {
		panic("Invalid key.")
	}
	if k > len(t.tables) {
		for len(t.tables) < k {
			t.tables = append(t.tables, nil)
		}
	}
	var tmp T = v
	t.tables[k-1] = &tmp
}

func main() {
	var t DATable[string]
	t.Insert("hello", 10)
	fmt.Println(t.Search(9))
}
