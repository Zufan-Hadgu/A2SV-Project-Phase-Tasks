package main

import (
    "fmt"
    "strings"
)

func reverse(s string) string {
    letters := []rune(s)
    for i, j := 0, len(letters)-1; i < j; i, j = i+1, j-1 {
        letters[i], letters[j] = letters[j], letters[i]
    }
    return string(letters)
}

func main() {
    var word string
    fmt.Print("Enter the string: ")
    fmt.Scanln(&word)

    word = strings.ToLower(word) 
    reversed := reverse(word)

    if reversed == word {
        fmt.Println("Palindrome")
    } else {
        fmt.Println("Not a Palindrome")
    }
}