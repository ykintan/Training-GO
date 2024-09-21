package main

import (
	"fmt"
	"time"
)

func RunHelloWorld() {
	println("Hello World")
}

func main() {
	go RunHelloWorld()
	fmt.Println("Main function")

	time.Sleep(1 * time.Second)
}
