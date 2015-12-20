package main

import (
	"fmt"
    "strconv"
    "bytes"
)

func main() {
	input := "1113122113"

	for r := 0; r < 50; r++ {
		
		var currentChar byte
		currentCharCount := 0
        var buffer bytes.Buffer
		
        for i:=0; i<len(input); i++ {
			currentChar = input[i]
			currentCharCount++
					
			if i == len(input)-1 || (input[i] != input[i+1]) {
                buffer.WriteString(strconv.Itoa(currentCharCount))
                buffer.WriteString(string(currentChar))
				currentCharCount = 0
			}
		}
		
        input = buffer.String()
        // fmt.Printf("%d: %d \n", r, buffer.Len())
	}	
	
	fmt.Println("length ", len(input))
}