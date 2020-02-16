package main 

import "fmt"

type student struct {
	name string
	age int
	email string
}

func (std student) introduce(){
	fmt.Println("hello my name is ", std.name)
}

type pupil struct {
	address string
	std student
}

func main(){
	hook := student{name: "hook"}
	pup := pupil{std: hook}
	pup.std.introduce()
}
