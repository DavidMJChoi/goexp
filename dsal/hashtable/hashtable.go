package main

import "fmt"

type DATable[T any] struct {
	tables []*T
}

func (t *DATable[T]) Search(k int) (T, bool) {
	if k < 0 || k >= len(t.tables) {
		// return zero value
		var zero T
		return zero, false
	}

	if t.tables[k] == nil {
		var zero T
		return zero, false
	}

	return *t.tables[k], true
}

func (t *DATable[T]) Insert(v T, k int) {
	if k < 0 {
		panic("Invalid key.")
	}
	for len(t.tables) < k {
		t.tables = append(t.tables, nil)
	}
	var tmp T = v
	t.tables[k-1] = &tmp
}

func (t *DATable[T]) Delete(k int) {
	if k >= 0 && k < len(t.tables) {
		t.tables[k] = nil
	}
}

func main() {
	var t DATable[string]
	t.Insert("hello", 10)
	t.Insert("hello", 9)
	t.Insert("hell", 8)
	t.Insert("hell00", 8)

	t.Delete(8)

	for i := range len(t.tables) {
		v, _ := t.Search(i)
		fmt.Println(i, ":", v)
	}
}
