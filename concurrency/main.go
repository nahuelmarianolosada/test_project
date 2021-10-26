package main

import (
	"fmt"
	"time"
)

func main() {
	//TestRoutine()
	testChannels()
}

func TestRoutine() {
	for i := 0; i < 10; i++ {
		go func(value int) {
			fmt.Printf("Hola desde la go routine %d \n", value)
		}(i)
		//go goroutine(i)
		//time.Sleep(100 * time.Millisecond)
	}

	time.Sleep(1000 * time.Millisecond)

	fmt.Println("Hola vithe")
}

func goroutine(value int) {
	fmt.Printf("Hola desde la go routine %d \n", value)
}

func testChannels() {
	c := make(chan string)

	go func() {
		c <- "Hello From Channel"
	}()

	//res := <-c

	fmt.Println(<-c)

}
