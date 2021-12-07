package main

import "fmt"

func test() {
	a := 1
	a++
	fmt.Println("test", "aaa")
}

func test1() {
	fmt.Println("test1")
}

func main() {
	a := 1
	fmt.Println("vim-go", a, "hello")
}
