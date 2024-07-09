package main

import (
	"fmt"
	"log"

	"github.com/nothing2512/gowes/client"
	"github.com/nothing2512/gowes/server"
)

func main() {
	runServer()
	client1()
	client2("cfce2309-1026-462b-a13b-f76614cfec20")
}

func client1() {
	c, err := client.Init("0.0.0.0:8080")
	if err != nil {
		log.Fatal(err)
		return
	}
	err = c.Connect()
	if err != nil {
		log.Fatal(err)
		return
	}

	c.Command(client.Message{
		Command: "cmd",
		Message: "msg",
	})

	c.OnMessage(func(m client.Message) {
		fmt.Println(m.Message)
	})
}

func client2(uid string) {
	c, err := client.Init("0.0.0.0:8080")
	if err != nil {
		log.Fatal(err)
		return
	}
	err = c.Connect()
	if err != nil {
		log.Fatal(err)
		return
	}

	c.Send(uid, client.Message{Message: "Hello"})
}

func runServer() {
	s := server.Init("11111111113333333333332222222222", "1112223334445556")
	s.OnCommand(func(m server.Message) {
		fmt.Println(m.Command, m.Message)
	})
	s.Start("0.0.0.0:8080")
}
