package main

import (
	"fmt"
	"reflect"
)

func main() {
	s := "Hello World"
	t := "你好，世界"

	tmp := s
	s = t
	t = tmp

	re := []byte(s)

	c := '世'

	fmt.Println(reflect.TypeOf(c))

	fmt.Println(s, t, re)

	fmt.Println("\xE4\xBD\xA0\xE5\xA5\xBD\xEF\xBC\x8C\xE4\xB8\x96\xE7\x95\x8C")
}
