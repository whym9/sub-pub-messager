package main

import (
	"fmt"

	"time"

	zmq "github.com/pebbe/zmq4"
)

func main() {

	x := zmq.PUB
	socket, _ := zmq.NewSocket(x)

	defer socket.Close()
	socket.Bind("tcp://*:5556")

	fmt.Println("Sending messages on port 5556")
	num := 0

	for {

		msg := fmt.Sprintf("message %v", num)
		num++

		socket.Send(msg, 0)
		time.Sleep(500 * time.Millisecond)
	}
}
