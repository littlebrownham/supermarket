package main

import (
	"fmt"

	"github.com/littlebrownham/supermarket/server"
)

func main() {
	server := server.New()
	err := server.Start()
	if err != nil {
		fmt.Printf("unable to start up %v", err)
	}
}
