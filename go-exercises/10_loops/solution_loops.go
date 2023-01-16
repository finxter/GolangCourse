// Exercise loops:

/*
 1. Complete the function divisbleByThree(n) to print the numbers that are
    divisible by 3. Ensure to print only upto a maximum of first 10 numbers and
    0 should be ignored.

 2. Complete the function printOddNums(start,end) to print the odd numbers between and including
    start and end.
*/
package main

import "fmt"

func divisbleByThree(n int) {
	count := 0
	fmt.Println("divisble by 3:")
	for i := 1; i <= n; i++ {
		if count == 10 {
			break
		}
		if i%3 == 0 {
			count++
			fmt.Println(i)
		}
	}
}

func printOddNums(start, end int) {
	fmt.Println("odd nos:")
	for i := start; i <= end; i++ {
		if i&1 == 1 { // 1, 3, 5, 9...
			fmt.Println(i)
		}
	}
}

func main() {
	divisbleByThree(25) // should print: 3,6,9,12,15,18,21,24
	divisbleByThree(45) // should print: 3,6,9,12,15,18,21,24,27,30
	divisbleByThree(0)  // nothing
	divisbleByThree(6)  // should print : 3,6

	printOddNums(7, 27) // 7, 9, 11, 13, 15, 17, 19, 21, 23, 25, 27
	printOddNums(2, 20) // 3,5,7,9,11,13,15,17,19
}
