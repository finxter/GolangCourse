// Exercise
// Arrays

/*
Create a to-do list to hold atmost 4 items. The items must include
the name of the task and it status.

Insert first three items (initially it is false) with status as
true.
Print the following to the console:
1. Items with status as true
2. total number of completed items in the to-do list

Update the fourth item with a new task.
Print the same info again.

Update the fourth item with status as true.
Print the same info again.

Hint: create a struct which holds name of the task and status, then
create an array of struct.
*/
package main

import "fmt"

type Todo struct {
	task   string
	status bool
}

// Print to the console
// 1. Items with status as true
// 2. total number of completed items in the to-do list
func printInfo(list [4]Todo) {
	var totalItems int
	fmt.Println("Items with status as true:")

	for i := 0; i < len(list); i++ {

		if list[i].status == true {
			fmt.Println(list[i].task, list[i].status)
			totalItems += 1
		}
	}
	fmt.Println("total number of completed items in the to-do list:", totalItems)
	fmt.Println("---------------------------------------")
}

func main() {

	// Create the to-do List here. You can create any tasks that
	// you do on a daily basis such as washing, cleaning etc.
	// Initialize only three items at the beginning with status as false.

	todoList := [4]Todo{
		{"Clean Room", false},
		{"Washing Machine", false},
		{"Cycling", false},
	}

	//  Insert first three items with status as true.
	todoList[0].status = true
	todoList[1].status = true
	todoList[2].status = true

	printInfo(todoList)

	// Insert fourth item status as false.
	todoList[3].task = "Prepare lunch"
	todoList[3].status = false
	printInfo(todoList)

	// Insert fourth item status as true.
	todoList[3].status = true
	printInfo(todoList)

}
