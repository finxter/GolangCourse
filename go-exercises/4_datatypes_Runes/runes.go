// Exercise
// 1. Replace the myString (currently "अ4Ω𐖸") with any other four chars of different lengths ranging from 1 to 4 bytes
// 2. Also print these individual char lengths by changing the chars with your own chars
// Run the program and see the output.
// For different chars refer : https://www.charset.org/utf-8

package main

import (
	"fmt"
	"reflect"
)

func main() {

	myString := "अ4Ω𐖸"         // change here
	fmt.Println(len(myString)) // for these chars length = 10(3 bytes + 1 byte + 2 bytes + 4 bytes)
	myRunes := []rune(myString)
	fmt.Println("string, runes:", myString, myRunes)

	fmt.Println("-----------Rune Unicode of each char in decimal and hex-----------------")
	fmt.Printf("%d, %x\n", myRunes[0], myRunes[0])
	fmt.Printf("%d, %x\n", myRunes[1], myRunes[1])
	fmt.Printf("%d, %x\n", myRunes[2], myRunes[2])
	fmt.Printf("%d, %x\n", myRunes[3], myRunes[3])

	fmt.Println("-----------Rune byte lengths-----------------")
	fmt.Println(reflect.TypeOf(myRunes)) // rune is internally int32
	fmt.Println(len("अ"))                // change here
	fmt.Println(len("4"))                // change here
	fmt.Println(len("Ω"))                // change here
	fmt.Println(len("𐖸"))                // change here

	fmt.Println("-----------Rune and byte-----------------")
	var r = 'a'      // default is rune
	var b byte = 'a' // explicit declaration
	fmt.Println("Type of r and b:", reflect.TypeOf(r), reflect.TypeOf(b))
}
