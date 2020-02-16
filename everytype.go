package main

import "fmt"

func hello(t interface{}){
	fmt.Printf("hello %t \n", t)
}

func hi(t ...interface{}) {
	fmt.Print("Hi")
	for _, v := range t {
		fmt.Printf(" %T,", v )
	}
	fmt.Println()
}

func main(){
	hello(3.123123213123)
	hello(true)
	hello("Hook")
	hi("nook", false, 10, 10e15)
}