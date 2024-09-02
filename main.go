package main

import "fmt"

func hello(msg string) {
	fmt.Println("Hello,", msg)
}

func main() {
	hello("Vanya")
	hello("Andrey")
	hello("Alexey")
	hello("Maxim")
	hello("Vasya")
}
