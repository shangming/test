package main

import "fmt"

func test() {
	a := 1
	a++
	fmt.Println("test", "aaaccc")
}

func test1() {
	fmt.Println("test1")
}

func main() {
	a := 1
	fmt.Println("vim-go", a, "hello")
	fmt.Println("vim-go", a, "hello")
}
