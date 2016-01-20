package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome!")
	go StartCamera()
	StartServer()
}
