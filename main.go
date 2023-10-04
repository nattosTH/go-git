package main

import "fmt"

func sum(a,b int) int {
	return a + b
}

func mul(a,b int) int {
	return  a *b
}

func display(msg string)  {
	fmt.Println(msg)
}

func main() {
	fmt.Println("hello Github")

	fmt.Println(sum(1,2))
	fmt.Println(mul(1,2))

	display("hello gitflow")
}