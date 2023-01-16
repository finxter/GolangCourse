// Range - used in loops to make life much easier for the developer

package main

import "fmt"

func main() {

	slice := []string{"learning", "golang", "!"}

	// regular for loop
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	fmt.Println("")

	// using ranges
	for i, element := range slice {
		fmt.Println(i, element)

	}

	// printing each char using for loop (internally each char is stored as rune
	// in go and can occupy more than 1 byte)
	fmt.Println("Using regular for loop to print each char ")

	for i := 0; i < len(slice[0]); i++ {
		fmt.Printf("%q", slice[0][i]) // trying to print each char of first string - "learning"
	}

	fmt.Println("")

	fmt.Println("Using range based for loop to print each char ")

	for i, element := range slice {
		fmt.Println(i, element, ":")
		for _, ch := range element {
			fmt.Printf("%q\n", ch)
		}
		break
	}
}




