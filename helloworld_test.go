package main

import (
	"fmt"
	"testing"
)

func Test_Say(t *testing.T) {
	msg := "Hi"
	Say(msg)
	fmt.Println("Test")
}

func Test_Say_Error(t *testing.T) {
	msg := "Hi"
	Say_Error(msg)
	fmt.Println("Test")
}
