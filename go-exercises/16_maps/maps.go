// Exercise
// maps

/*
   // add prints wherever necessary
   1. Create an empty map with key as string and value as struct
      (Vertex).The Vertex struct is given which has latitude and longitude.
      Insert keys with company names - google and apple. Update
      them with lat,long as (37.422131, -122.084801) for google
      and (37.33182,-122.03118).

	2. Given a slice of strings, calculate the word count. Print the
	   word count
	   Hint: use range for loop to go through the slice and use
	   the map check for key existance for word count.
*/

package main

import "fmt"

type Vertex struct {
	lat, long float64
}

// 1
func company_Vertex() {
	fmt.Println("hello") // write code here

}

// 2
func word_count(s []string) {
	fmt.Println("hello") // write code here
}
func main() {

	company_Vertex()

	slice := []string{
		"aa",
		"bbb",
		"c",
		"aa",
		"bbb",
		"bbb",
		"bbb",
		"aa",
	}
	word_count(slice) // must print map[aa:3 bbb:4 c:1]

}
