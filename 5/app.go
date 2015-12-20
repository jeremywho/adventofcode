package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	
	b, _ := ioutil.ReadFile("input.txt")
	input := string(b[:])	
	lines := strings.Split(input, "\n") 
		
	goodLineCount := 0
	
	for _, line := range lines {
		fmt.Println(line)
		fmt.Println(hasDoubleLetters(line))
		if  hasDoubleLetters(line) &&
			!hasBadString(line) &&
			numVowels(line) > 2 {
				goodLineCount++
			}
	}	
	fmt.Println("Number of nice strings: ", goodLineCount)
}

func hasDoubleLetters(s string) bool {
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			return true
		}
	}
	return false;
}

func hasBadString(s string) bool {
	//this is not efficient, but it is easy to read
	return 	strings.Contains(s, "ab") ||
			strings.Contains(s, "cd") ||
			strings.Contains(s, "pq") ||
			strings.Contains(s, "xy")
}

func numVowels(s string) int {
	n := 0
	for i := 0; i < len(s); i++ {
		if  s[i] == 'a' ||
			s[i] == 'e' ||
			s[i] == 'i' ||
			s[i] == 'o' ||
			s[i] == 'u' {
			n++
		}
	}
	return n;
}