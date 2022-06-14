package main

import "fmt"

func main() {
	name := "bob"
	fmt.Println("Hello World")
	fmt.Printf("The name is %v\n", name)
	capture := fmt.Sprintf("string capture")
	fmt.Println(capture)
}
