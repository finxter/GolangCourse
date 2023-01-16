// Exercise
// Pointers

/*
	1. Declare a function which accepts an int pointer and changes the value passed to it
	   from 25 to 50

	2. For the function changeGrade() which accepts a slice of struct pointer, updates
	the grade of student at index 2 to 89 from 85.

	Add suitable print statements whereever necessary to observe change in values.
*/

package main

import "fmt"

type studentGrades struct {
	name  string
	grade int
}

// 2
func changeGrade(s *[]studentGrades) {

}

func main() {

	students := []studentGrades{
		{
			"Tom",
			75,
		},
		{
			"Ram",
			99,
		},
		{
			"Jim",
			85,
		},
		{
			"Krishna",
			95,
		},
	}
	fmt.Println(students)

}
