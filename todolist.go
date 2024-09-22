package main

import "fmt"

func addtodo(){

}

func removetodo(){

}

func starttodolist(){

}

func main(){
	//declares some variables
	var todolist = [10] string{}
	var maketodo string
	//asking if the user creats a todo
	fmt.Println("would you like to make a todolist? y/n")
	fmt.Scan(&maketodo)
	//yes or no to the todo
	//the yes outcome
	if maketodo == "y" {
		var action string
		fmt.Println("would u like to see ur todo,add element,remove element")
		fmt.Scan(&action)
		//see todo
		if action == "seetodo" {
			fmt.Print(todolist)
		//add todo elemnt
		}else if action == "addelement" {
			addtodo()
		//remove todo elemnt
		}else if action == "removeelement" {
			removetodo()
			
		}else {fmt.Println("invalid command")}

	//the no outcome
	} else if maketodo == "n" {
		fmt.Println("have a nice day")
	//incase user provided something else	
	} else {fmt.Println("invalid letter,insert y or n")}

}