package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	
	b, _ := ioutil.ReadFile("input.txt")
	input := string(b[:])
	
	x := 0
	y := 0
	
	roboX := 0
	roboY := 0
	
	santaX := 0
	santaY := 0
	
	m := make(map[string]int)
	m["0,0"] = 1
	
	robo := make(map[string]int)
	robo["0,0"] = 1
	
	santa := make(map[string]int)
	santa["0,0"] = 1
		
	for i := 0; i < len(input); i++ {
		if input[i] == '>' {
			x++
			if (i % 2) == 0 {
				santaX++
			} else {
				roboX++
			}
		} else if input[i] == '<' {
			x--
			if (i % 2) == 0 {
				santaX--
			} else {
				roboX--
			}
		} else if input[i] == '^' {
			y++
			if (i % 2) == 0 {
				santaY++
			} else {
				roboY++
			}
		} else if input[i] == 'v' {
			y--
			if (i % 2) == 0 {
				santaY--
			} else {
				roboY--
			}
		} 
		
		coord := fmt.Sprintf("%d,%d", x, y)
		m[coord] = 1
		
		coord = fmt.Sprintf("%d,%d", santaX, santaY)
		santa[coord] = 1
		
		coord = fmt.Sprintf("%d,%d", roboX, roboY)
		robo[coord] = 1	
	}
	
	//merge santa and robo santa
	result := 0
	for i := -999; i < 1000; i++ {
		for j := -999; j < 1000; j++ {
			
			santaVisited := false
			roboVisitied := false
			
			coord := fmt.Sprintf("%d,%d", i, j)
			_, santaVisited = santa[coord]
			_, roboVisitied = robo[coord]
			
			if santaVisited || roboVisitied {
				result++
			}
		}
	}
	
	fmt.Printf("\nPart 1\n")
	fmt.Println(len(m))
	fmt.Printf("\nPart 2\n")
	fmt.Println(result)
}