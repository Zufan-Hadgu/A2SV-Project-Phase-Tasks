package main

import (
	"fmt"
	"strings"
)

func main(){
	var sentence string
	fmt.Print("Enter the String")
	fmt.Scanln(&sentence)

	lower := strings.ToLower(sentence)
	words := strings.Fields(lower)
	count := make(map[string]int)
	for _,word := range words{
		count[word] += 1
	}
	fmt.Print(count)

}