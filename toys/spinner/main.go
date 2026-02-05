package main

import (
	"bytes"
	"fmt"
	"time"
)

func fib(n int) (r int) {
	if n == 1 || n == 2 {
		return 1
	} else {
		return fib(n-1) + fib(n-2)
	}
}

func spinner() {
	spinner := `-\|/`
	for {
		for _, c := range spinner {
			fmt.Printf("\r%c", c)
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func pointCounter() {
	buf := bytes.NewBuffer(make([]byte, 3))
	for {
		for range 3 {
			buf.WriteByte('.')
			fmt.Printf("\r\033[K%s", buf.String())
			time.Sleep(100 * time.Millisecond)
		}
		buf.Reset()
	}
}

func main() {
	// go spinner()

	go pointCounter()
	fmt.Printf("\r%v\n", fib(45))

	// time.Sleep(5 * time.Second)
	// fmt.Println("Done!")
}
