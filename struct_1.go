package main

import "fmt"

type student struct {
	name string 
	age int 
	email string
}

func main(){
	var a student
	a.name = "Hook"
	a.age = 18 
	a.email = "HooK@Hook.Hk"

	b := student{"nook", 3 , "nook@nook.nk"}

	c := student{name: "mook", email: "mook@mook.mk"}
	
	d := student{age: 20}

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}