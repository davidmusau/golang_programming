package main

import (
	"fmt"
	"time"
)

func pinger(c chan string) {

	for i := 0; ; i++ {

		c <- "ping"
	}
}
func printer(c chan string) {
	const UnixDate string = "Mon Jan _2 15:04:05 MST 2006"

	dt := time.Now()
	ts := time.Second * 1
	for {
		msg := <-c
		fmt.Println(msg, dt.Format(time.UnixDate))
		time.Sleep(ts)
	}
}

func ponger(c chan string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}

func main() {
	//initialize channel
	var c chan string = make(chan string)

	go pinger(c)
	go printer(c)
	go ponger(c)
	//offer keyboard interruptions
	var input string
	fmt.Scanln(&input)
}
