package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {
	str := "hello world"

	fmt.Println(str[0:1])       // prints h
	fmt.Println(str[0])         // prints 104; ascii value
	fmt.Println(string(str[0])) // prints h; string converts ascii
	fmt.Println("hello"[1])     // prints 101; ascci for 'e'

	fmt.Println(strings.ToUpper("hello"))

	splitStr := strings.Split(str, "")
	fmt.Println(str)
	fmt.Println(splitStr)
	fmt.Printf("%#v\n", splitStr)

	fmt.Println(reflect.TypeOf(splitStr))

	fmt.Println(strings.Contains("pdog", "dog"))
}
