package main

import "fmt"

//function to add a todo
func addtodo(todolist *[10]string){
	for i := 0; i < 10; i++ {
		fmt.Println("enter todo: ")
		fmt.Scan(&todolist[i])
	}

}

//function to remove a todo
func removetodo(todolist *[10]string){
	var n int
	fmt.Println("enter todo number that u want to be removed")
	fmt.Scan(&n)
	todolist[n-1] = ""

}

//function to print all of the todolist
func seetodo(todolist *[10]string){
	for i := 0; i < 10; i++ {
		fmt.Println(todolist[i])
	}
}

func main(){
	//declaring some variables
	var quit string
	var todolist = [10] string{}
	var maketodo string

	//asking if the user creats a todo
	fmt.Println("would you like to make a todolist? y/n")
	fmt.Scan(&maketodo)

	//option of yes
	if maketodo == "y" {
		var action string
	for  {
		fmt.Println("would u like to seetodo,addtodo,removetodo")
		fmt.Scan(&action)
		
		//see todo
		if action == "seetodo" {
			seetodo(&todolist)

		//add todo elemnt
		}else if action == "addtodo" {
			addtodo(&todolist)

		//remove todo elemnt
		}else if action == "removetodo" {
			removetodo(&todolist)

		//incase user adds a diffrent command
		}else {fmt.Println("invalid command")}

		fmt.Println("quit? y/n")
		fmt.Scan(&quit)
		if quit == "y" {
			break
		}else if quit == "n" {
			
		}else{fmt.Println("command invalid")}
	}	

	//option of no
	} else if maketodo == "n" {
		fmt.Println("have a nice day")
	//incase user provided something else	
	} else {fmt.Println("invalid letter,insert y or n")}
	
}