package main

import (
	"fmt"
	"github.com/satori/go.uuid"
	"hello/world"
)

func main() {
	fmt.Println("Hello World")

	u1 := uuid.Must(uuid.NewV4(), nil)
	fmt.Printf("UUIDv4: %s\n", u1)

	world.Say()

	world.You()

	world.Tab()
	world.CAY()
}
