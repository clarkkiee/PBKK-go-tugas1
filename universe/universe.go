package main

import "fmt"

func main () {

	var input int

	fmt.Print("Enter a number : ")
	fmt.Scanln(&input)

	if input == 42 {
		fmt.Println("Universe")
	} else {
		fmt.Println(input)
	}

}