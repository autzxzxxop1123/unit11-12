package main

import "fmt"

type I interface {}
	
func main(){ 
	var i I
	i = "hello"
	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	b := i.(bool)
	fmt.Print(b)
}