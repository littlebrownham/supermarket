package main

import (
	"fmt"

	"github.com/littlebrownham/supermarket/server"
)

func main() {
	server, err := server.New()
	if err != nil {
		fmt.Println("unable to start up")
	}
	err = server.Start()
	if err != nil {
		fmt.Println("unable to start up")
	}
}
