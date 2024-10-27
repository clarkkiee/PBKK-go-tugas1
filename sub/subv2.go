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
	
	input = strings.TrimSpace(input)
	var chunked []string
	
	temp := strings.Split(input, " ")
	if len(temp) < 3 {
		fmt.Println("Minimum length of string is 3 words")
		os.Exit(1)
	}
	for i := range len(temp) {
		chunked = append(chunked, reverse(temp[i]))
	} 

	for _, chunk := range chunked {
		fmt.Println(chunk)
	}

}