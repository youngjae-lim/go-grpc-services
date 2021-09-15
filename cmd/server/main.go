package main

import (
	"fmt"
	"log"
)

// Run initializes and starts our gRPC server
func Run() error {
	fmt.Println("Rocket Service Starting...")
	return nil
}

func main() {
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}