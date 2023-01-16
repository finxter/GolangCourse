// Exercise variables:
// Print all the variables to the terminal

/*
 1. Declare variable for your favorite movie and print it.
 2. Using Block declaration use two variables to store your name and age.
 3. Declare a variable for next leap year but do not assign it. Assign a value to it on the next line.
 4. Using short assignment declare two variables to hold the winner of golden boot in FIFA 2022 and his T-shirt number.
 (e.g. Messi with number 10)
*/

package main

import "fmt"

func main() {
	// 1
	var myFavouriteMovie string = "Batman"
	fmt.Println("movie:", myFavouriteMovie)

	// 2
	var (
		name string = "test"
		age  int    = 35
	)
	fmt.Println("name, age:", name, age)

	// 3
	var nextLeapYear string
	nextLeapYear = "2024"
	fmt.Println("next leap year:", nextLeapYear)

	// 4
	goldenBootWinner, shirtNumber := "Mbappe", 10
	fmt.Println("golden boot winner:", goldenBootWinner, shirtNumber)
}
