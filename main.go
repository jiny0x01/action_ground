package main

import "fmt"

func makeConflict() {
	fmt.Println("This is occur go fmt error")
}

func main() {
	fmt.Println("Hello Golang!")
	fmt.Println("Hello Golang!")
	fmt.Println("Hello Golang!")
}
