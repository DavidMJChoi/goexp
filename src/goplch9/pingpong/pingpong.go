package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

var ch = make(chan struct{})

// var mu sync.Mutex

func main() {
	pingpongCnt := int64(0)

	go func() {
		for {
			ch <- struct{}{}
			<-ch

			// mu.Lock()
			// pingpongCnt++
			// mu.Unlock()
			atomic.AddInt64(&pingpongCnt, 1)
		}
	}()

	go func() {
		for {
			<-ch

			atomic.AddInt64(&pingpongCnt, 1)
			// mu.Lock()
			// pingpongCnt++
			// mu.Unlock()

			ch <- struct{}{}
		}
	}()

	time.Sleep(time.Second)
	fmt.Println(pingpongCnt)
}
