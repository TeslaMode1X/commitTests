package main

import "fmt"

func hello(msg string) {
	fmt.Println("Hello,", msg)
}

func loopThatMakesSense() {
	i := 0
	for i++; i < 5; i++ {
		fmt.Println(i)
	}
}

func main() {
	hello("Vanya")
	hello("Andrey")
	hello("Alexey")
	hello("Maxim")
	hello("Vasya")

	// for fun
	loopThatMakesSense()
}
