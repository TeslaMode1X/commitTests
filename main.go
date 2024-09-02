package main

import "fmt"

func hello(msg string) {
	fmt.Println("Hello,", msg)
}

func infinityLoop() {
	i := 0
	for {
		i++
	}
}

func main() {
	hello("Vanya")
	hello("Andrey")
	hello("Alexey")
	hello("Maxim")
	hello("Vasya")

	// for fun
	infinityLoop()
}
