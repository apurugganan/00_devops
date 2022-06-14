package main

import (
	"fmt"
	"strings"
)

func main() {
	name := "bob"
	fmt.Println("Hello World")
	fmt.Printf("The name is %v\n", name)

	capture := fmt.Sprintf("string capture")
	fmt.Println(capture[0:1])       // prints s
	fmt.Println(capture[0])         // prints 115; ascii value
	fmt.Println(string(capture[0])) // prints s; string converts ascii
	fmt.Println("hello"[1])         // prints 101

	fmt.Println(strings.ToUpper("hello"))

}
