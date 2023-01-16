// Exercise if-else:
// Complete the below function.

/*
 1. Complete the function such that if the string is "Finxter" or
    "finxter" then print success, else print failure

 2. Complete the function given such that below conditions are met
    a. If the number is divisible by 3, print number divisible by 3
    b. If the number is divisible by 5, print number divisible by 5
    c. If the number is divisible by both 3 and 5, then
    i.) if the number is >=45 and <=100, print the number
    ii.) otherwise print number + 1 (ie. the next number)
    d.) If it is neither divisible by 3 or 5, then print the message "not a good number"
    e.) If the number is 0 or 1, print the message "retry without 0 and 1".

    Use return statement wherever possible to make it efficient.
*/
package main

import "fmt"

// 1

func areYouAFinxter(s string) {
	if s == "Finxter" || s == "finxter" {
		fmt.Println("success")
	} else {
		fmt.Println("failure")
	}
}

//2

func checkForThreeAndFive(number int) {
	if number == 0 || number == 1 {
		fmt.Println("retry without 0 and 1")
		return
	} else {
		if number%3 == 0 && number%5 == 0 {
			if number >= 45 && number <= 100 {
				fmt.Println("number:", number)
			} else {
				fmt.Println("next number:", number+1)
			}
		} else if number%3 == 0 {
			fmt.Println("number divisble by 3")
		} else if number%5 == 0 {
			fmt.Println("number divisble by 5")
		} else {
			fmt.Println("not a good number")
		}

	}

}

func main() {
	// test 1
	areYouAFinxter("finxter")    // should print success
	areYouAFinxter("notfinxter") // should print failure
	areYouAFinxter("Finxter")    // should print success

	// test 2
	checkForThreeAndFive(9)   // should print number divisble by 3
	checkForThreeAndFive(25)  // should print number divisble by 5
	checkForThreeAndFive(15)  // should print next number: 16
	checkForThreeAndFive(60)  // should print number: 60
	checkForThreeAndFive(45)  // should print number: 45
	checkForThreeAndFive(100) // should print number divisble by 5
	checkForThreeAndFive(1)   // should print retry without 0 and 1
	checkForThreeAndFive(0)   // should print retry without 0 and 1
	checkForThreeAndFive(17)  // should print not a good number

}
