	package main 

	import (
		"fmt"
		"bufio"
		"os"
		"strings"
	)

	func reverse (s string) string {
		rns := []rune(s)
		for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
			rns[i], rns[j] = rns[j], rns[i]
		}
		return string(rns)
	}

	func main () {
			
		read := bufio.NewReader(os.Stdin)
		fmt.Print("Input a string : ")
		input, _ := read.ReadString('\n')

		for _, word := range strings.Fields(input) {
			if len(strings.Fields(input)) < 3 {
				fmt.Println("Minimum length of string is 3 words")
				break
			} 
			fmt.Println(reverse(word))
		}

	}