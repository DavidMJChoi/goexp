package main

import (
	"fmt"
	"unsafe"
)

type hmap struct {
	count      int
	B          uint8
	buckets    unsafe.Pointer
	oldbuckets unsafe.Pointer
	// ...
}

func main() {
	var m map[int]int

	fmt.Println(len(m))
}
