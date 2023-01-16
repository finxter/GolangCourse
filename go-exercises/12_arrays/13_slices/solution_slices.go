// Exercise

/*
 1. Complete the function slice_index().
    Declare a slice with its initial array of six integers and initialize it from 1 to 6.
    i.) Create a slice from index 2 upto 5.
    ii.) Create a slice from index 3 until end of array.
    Print the new slices.

 2. Complete the merge() function
    Two slices are given. Merge slice2 to slice1 and print.

 3. Complete the preallocate() function.
    Preallocate a slice with length 4 to hold flowers. Update it with the flowers - rose, lotus and tulips

 4. A slice is passed as param to a function slice_params(). In the function print the elements using for loop
    and then update index 1 and index 2 with new values - (44,55). Print the slice in the main function
    and observe if it also changes.Can you explain why the values of the slice in main() also changed ?

    Hint: slice is a reference type.
*/
package main

import "fmt"

// 1
func slice_index() {
	mySlice := []int{1, 2, 3, 4, 5, 6}
	slice1 := mySlice[2:5]
	slice2 := mySlice[3:]
	fmt.Println(slice1, slice2)
}

// 2
func merge() {
	slice1 := []int{1, 3, 5}
	slice2 := []int{2, 4, 6}
	combine := append(slice1, slice2...)
	fmt.Println("combined:", combine)
}

// 3

func preallocate() {
	flowers := make([]string, 4)

	flowers[0] = "rose"
	flowers[1] = "lotus"
	flowers[2] = "tulips"
	fmt.Println(flowers)
}

// 4
func slice_params(slice []int) {
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}
	slice[1] = 44
	slice[2] = 55
}
func main() {

	slice_index() //[3 4 5] [4 5 6]

	merge() // [1 3 5 2 4 6]

	preallocate() // [rose lotus tulips ]

	nums := []int{3, 4, 5, 6, 7, 8, 9}
	slice_params(nums)

	fmt.Println(nums) //  [3 44 55 6 7 8 9]

}
