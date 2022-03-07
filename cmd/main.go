package main

import "fmt"

func main() {
	hello("Hello, World!")
}

func hello(s string) {
	msg := s
	fmt.Println(msg)
}
