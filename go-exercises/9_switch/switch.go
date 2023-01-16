package main

import "fmt"

// Exercise Switch:

// Complete the below functions by adding a switch statement

/*
1. The function checkWeekOfTheDay takes a day as input and checks if the day is
	a. If day is between 1 - 7, then print First Week
	b. If day is between 8 - 14, then print Second week
	c. If day is between 15 - 21, then print Third week
	d. If day is between 22 - 31, then print fourth week
	e. For all other days print - not a valid day
	Hint: Try using case list

2. The function getDesignation takes team member's userId, print their designations
		userId 0 : Product Head, Developer and Tester
		userId 1:  Developer and Tester
		userId 2 : Tester
		userId 3 : QA
		any other value: not a valid team member
	Hint: Use fallthrough when you need to have multiple designations
*/

// 1
func checkWeekOfTheDay(day int) {
	fmt.Println("hello") // Write code here
}

// 2
func getDesignation(userId int) {
	fmt.Println("hello") // Write code here
}

func main() {
	checkWeekOfTheDay(3)  //output: First Week
	checkWeekOfTheDay(10) //output: Second Week
	checkWeekOfTheDay(16) //output: Third Week
	checkWeekOfTheDay(22) //output: Fourth Week
	checkWeekOfTheDay(28) //output: Fourth Week
	checkWeekOfTheDay(0)  //output: not a valid day
	checkWeekOfTheDay(32) //output: not a valid day

	// Expected Output
	getDesignation(0) // output:  Product Head,Developer,Tester
	getDesignation(1) // output:  Developer,Tester
	getDesignation(2) //output:  Tester
	getDesignation(3) //output:  QA
	getDesignation(6) //output:  Not a valid team member
}
