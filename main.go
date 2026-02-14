package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
)

func main() {
	var w io.Writer = os.Stderr
	fmt.Println(reflect.TypeOf(w))
}
