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
	 student
}

func main() {
	var pup pupil
	pup.name = "Hook"
	pup.introduce()

}