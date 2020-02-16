package main 

import "fmt"

type student struct {
	name string
	age int 
	email string 
}

func main(){
	std := student{name: "Hook"}
	p := &std
	(*p).age = 18
	p.email = "Hook@Hook.hk"

	fmt.Println(std)
}