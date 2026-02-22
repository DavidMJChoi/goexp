package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()

	fmt.Println("Commencing countdown.")

	tick := time.Tick(time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		select {
		case <-tick:
		case <-abort:
			fmt.Println("FUCK IS ABORTED!")
			return
		}
	}

	launch()
}

func launch() {
	fmt.Println("FUCK THIS WORLD")
}
