// Exercise

// Structs

/*
   Complete the functions given below with the following requirements.
   You are given four coordinates (x1,y1), (x2,y2), (x3,y3) and (x4,y4).
   Check if they form a square.
   e.g. 1
   p1 = { 10, 10 }, p2 = { 20, 10 }, p3 = { 10, 20 }, p4 = { 20, 20 }
   Ans: Yes
   e.g. 2
   p1 = { 10, 10 }, p2 = { 20, 20 }, p3 = { 10, 20 }, p4 = { 20, 20 }
   Ans:No

   Hint from examples: to form a square the horizontal distances or x-axis distances
   and vertical distances or y-axis distances must all be same.
   ie. p2 - p1 and p4 - p3 will give us horizontal or x-axis distances
       p3 - p1 and p4 - p2 will give us vertical or y-axis distances

   e.g. 1:
   Horizontal distances:
   d1 = p2 - p1 =  (20-10 + 10 -10) = 10
   d2 = p4 - p3 =  (20-10 + 20 -20) = 10
   As you can see no need of y distances here as it will be 0
   Vertical distances:
   d3 = p3 - p1 = (10-10 + 20 -10) = 10
   d4 = p4 - p2 = (20-20 + 20 -10) = 10
   As you can see no need of x distances here as it will be 0

   As d1 = d2 = d3 = d4, thus ans is Yes

   e.g. 2:
   Horizontal distances:
   d1 = p2 - p1 =  (20-10 + 20 -10) = 20
   d2 = p4 - p3 =  (20-10 + 20 -20) = 10

   As d1 != d2, it cant be a square, no need of any further calculations

*/

package main

import "fmt"

type Coordinates struct {
	x uint16
	y uint16
}

type Square struct {
	p1 Coordinates
	p2 Coordinates
	p3 Coordinates
	p4 Coordinates
}

// This function must set the square coordinates p1,p2,p3 and p4
func setSquareCoordinates(a, b, c, d Coordinates) Square {
	// Your code here
	return s
}

// d1 = p2 - p1
// d2 = p4 - p3
// d3 = p3 - p1
// d4 = p4 - p2
func checkIfSquare(s Square) bool {

	// Your code here

	return true // or  false
}

func main() {
	// Set square coordinates for
	// p1 = { 10, 10 }, p2 = { 20, 10 }, p3 = { 10, 20 }, p4 = { 20, 20 }
	square := setSquareCoordinates( /*Your code here*/ )
	fmt.Println(checkIfSquare(square)) // output: true

	// p1 = { 10, 10 }, p2 = { 20, 20 }, p3 = { 10, 20 }, p4 = { 20, 20 }
	square = setSquareCoordinates( /*Your code here*/ )
	fmt.Println(checkIfSquare(square)) // output: false

	// p1 = { 20, 20 }, p2 = { 40, 20 }, p3 = { 20, 40 }, p4 = { 30, 40 }
	square = setSquareCoordinates( /*Your code here*/ )
	fmt.Println(checkIfSquare(square)) // output: false

	// p1 = { 20, 20 }, p2 = { 40, 20 }, p3 = { 20, 40 }, p4 = {40, 40 }
	square = setSquareCoordinates( /*Your code here*/ )
	fmt.Println(checkIfSquare(square)) // output: true

}
