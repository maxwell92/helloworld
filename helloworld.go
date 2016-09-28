package main

import "fmt"

func Say(msg string) {
	fmt.Println(msg)
}

func main() {
	fmt.Println("Hello World")
	msg := "Hi"
	Say(msg)
}
