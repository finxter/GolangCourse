// Exercise Functions:
// Use meaningful function names
// For all functions, print the result by calling the functions from main.

/*
 1. Write a function which takes first and last names as input,
    returns a string by adding them with +.

 2. Write a function that takes three numbers as inputs and returns
    the sum of the three.

3. Write a function which returns a greeting message to the user.

 4. Write a function which takes a number and returns multiple values:
    number + 1,number + 2 and number + 3.
*/
package main

import "fmt"

// 1
func combineFirstAndLast(first string, last string) string {
	return (first + last)
}

// 2
func threeNumberSum(x, y, z int) int {
	return (x + y + z)
}

// 3
func greetings() string {
	return "Have a Good Day"
}

// 4
func numbers(n int) (int, int, int) {
	return n + 1, n + 2, n + 3
}
func main() {
	fmt.Println(combineFirstAndLast("Leo", "Messi"))
	fmt.Println(threeNumberSum(3, 4, 5))
	fmt.Println(greetings())
	fmt.Println(numbers(5))
}
