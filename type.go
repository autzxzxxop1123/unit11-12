package main

import "fmt"

func main(){
	var studentName [10]string
	var studentAge [10]int
	var studentEmail [10]string

	studentName[0] = "Hook"
	studentAge[0] = 18 
	studentEmail[0] = "Hook.12"

	fmt.Println( studentName[0] , studentAge[0], studentEmail[0])
}