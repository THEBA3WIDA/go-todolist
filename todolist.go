package main

import "fmt"

func addtodo(){

}

func removetodo(){

}

func starttodolist(){

}



func main(){
	//var todolist = [10] string{}
	var maketodo string
	fmt.Println("would you like to make a todolist? y/n")
	fmt.Scan(&maketodo)
	if maketodo == "y" {

		
		
		
	} else if maketodo == "n" {
		fmt.Println("have a nice day")
	} else {fmt.Println("invalid letter,insert y or n")}

}