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
