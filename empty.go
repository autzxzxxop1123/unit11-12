package main

import "fmt"

type I interface {}


func desc(i I) {
	fmt.Printf("%v, %T \n", i, i)
}

func main(){
	var i I
	i = 3.114214123
	desc(i)
	i = "hello world"
	desc(i)
}