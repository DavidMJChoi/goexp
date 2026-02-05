package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	conn1, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Println("unable to read from localhost:8080")
	}
	defer conn1.Close()
	clock1 := make(chan string)
	go func() {
		for {
			buf := make([]byte, 1024)
			n, _ := conn1.Read(buf)
			clock1 <- string(buf[:n])
		}
	}()

	conn2, err := net.Dial("tcp", "localhost:8081")
	if err != nil {
		log.Println("unable to read from localhost:8081")
	}
	defer conn2.Close()
	clock2 := make(chan string)
	go func() {
		for {
			buf := make([]byte, 1024)
			n, _ := conn2.Read(buf)
			clock2 <- string(buf[:n])
		}
	}()

	for {

		time1 := <-clock1
		time2 := <-clock2

		fmt.Print(time1, " ", time2)
		time.Sleep(1 * time.Second)
		fmt.Print("\033[2J\033[H")
	}

}
