package main

import "fmt"

func addtodo(todolist *[10]string){
	for i := 0; i < 10; i++ {
		fmt.Println("enter todo: ")
		fmt.Scan(&todolist[i])
	}

}

func removetodo(todolist *[10]string){
	var n int
	fmt.Println("enter todo number that u want to be removed")
	fmt.Scan(&n)
	todolist[n] = ""

}

func seetodo(todolist *[10]string){
	for i := 0; i < 10; i++ {
		fmt.Println(todolist[i])
	}
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
			seetodo(&todolist)
		//add todo elemnt
		}else if action == "addelement" {
			addtodo(&todolist)
		//remove todo elemnt
		}else if action == "removeelement" {
			removetodo(&todolist)
			
		}else {fmt.Println("invalid command")}

	//the no outcome
	} else if maketodo == "n" {
		fmt.Println("have a nice day")
	//incase user provided something else	
	} else {fmt.Println("invalid letter,insert y or n")}

}