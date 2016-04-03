// Channels provide a way for two goroutines to
// communicate wit a one another and synchronize 
// their execution

package main

import (
	"fmt"
	"time"
)

// We can specify a direction on a channel type thus
// restricting it to either sending or receiving
// Now 'c' can only be sent to
func pinger (c chan<- string) {
	for i := 0; ;i++ {
		// The <- (left arrow) operator is used to send and recive messages on the channel
		// This means send "ping"
		c <- "ping"
	}
}

// Here 'c' can only receive
func printer(c <-chan string) {
	for {
		// This means recieve a mesage and store it in msg
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func ponger (c chan string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}

func main(){
	// A channel type is represented with the keyword
	// 'chan' folloewd by the type of the things that
	// are passed on the channel
	var c chan string = make(chan string)
	go pinger(c)
	go ponger(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)
}
