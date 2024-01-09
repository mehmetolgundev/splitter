package main

import (
	"log"

	"github.com/mehmetolgundev/splitter/pkg/listener"
)

func main() {
	err := listener.NewTCPListener().Listen()
	if err != nil {
		log.Fatal(err)
	}
}
