package main

import "fmt"

const greetPrefix = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return greetPrefix + name
}

func main() {
	fmt.Println(Hello("World"))
}