package main

import (
	"fmt"
	"log"
	"net/rpc"
)

var name string = "Omar"

func main() {
	client, err := rpc.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Println(err)
		return
	}

	var reply string
	err = client.Call("Greeting.SayHello", name, &reply)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Printf("%s\n", reply)
}
