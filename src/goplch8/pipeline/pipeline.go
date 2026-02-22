package main

import "fmt"

func main() {
	cToS := make(chan int)
	sToP := make(chan int)
	N := 10

	go counter(N, cToS)
	go squarer(cToS, sToP)
	printer(sToP)
}

func counter(N int, cToS chan<- int) {
	for i := range N {
		cToS <- i
	}
	close(cToS)
}

func squarer(cToS <-chan int, sToP chan<- int) {
	for x := range cToS {
		sToP <- x * x
	}
	close(sToP)
}

func printer(sToP <-chan int) {
	for x := range sToP {
		fmt.Println(x)
	}
}
