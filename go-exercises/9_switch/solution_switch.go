package main

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
	Hint: Use fallthrough when you need to have both
*/

import "fmt"

// 1
func checkWeekOfTheDay(day int) {

	switch day {

	case 1, 2, 3, 4, 5, 6, 7:
		fmt.Println(" First Week")
	case 8, 9, 10, 11, 12, 13, 14:
		fmt.Println(" Second Week")
	case 15, 16, 17, 18, 19, 20, 21:
		fmt.Println(" Third Week")
	case 22, 23, 24, 25, 26, 27, 28, 29, 30, 31:
		fmt.Println(" Fourth Week")
	default:
		fmt.Println("not a valid day")
	}
}

// 2
func getDesignation(userId int) {

	switch userId {
	case 0:
		fmt.Println("Product Head")
		fallthrough
	case 1:
		fmt.Println("Developer")
		fallthrough
	case 2:
		fmt.Println("Tester")
	case 3:
		fmt.Println("QA")
	default:
		fmt.Println("Not a valid team member")
	}
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
