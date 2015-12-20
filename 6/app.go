package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
)

type handle func(string, map[string]int)

func main() {
	
	b, _ := ioutil.ReadFile("input.txt")
	input := string(b[:])	
	lines := strings.Split(input, "\n")
	
	board := make(map[string]int) 
	
	for _, line := range lines {		
		
		if strings.HasPrefix(line, "toggle") {
			handleLine(line, board, handleToggle)
		} else if strings.HasPrefix(line, "turn on") {
			handleLine(line, board, handleOn)
		} else if strings.HasPrefix(line, "turn off") {
			handleLine(line, board, handleOff)
		}
	}	
	//400410
	fmt.Println(len(board))
}

func handleOn(loc string, b map[string]int) {
	b[loc] = 1
}

func handleOff(loc string, b map[string]int) {
	delete(b, loc)
}

func handleToggle(loc string, b map[string]int) {
	_, ok := b[loc]
	if ok {
		delete(b, loc)
	} else {
		b[loc] = 1
	}
} 

func handleLine(line string, b map[string]int, h handle) {	
	x1, y1, x2, y2 := parseLine(line)
	for i := x1; i <= x2; i++ {
		for j := y1; j <= y2; j++ {
			loc := fmt.Sprintf("%d,%d", i, j)
			h(loc, b)
		}
	}
}

func parseLine(line string) (int,int,int,int) {
	t := strings.Split(line, " ")
	var p1 []string
	var p2 []string
	
	if t[0] == "toggle" {
		p1 = strings.Split(t[1], ",")
		p2 = strings.Split(t[3], ",")
	} else {
		p1 = strings.Split(t[2], ",")
		p2 = strings.Split(t[4], ",")
	}
	
	x1, _ := strconv.Atoi(p1[0])
	y1, _ := strconv.Atoi(p1[1])
	
	x2, _ := strconv.Atoi(p2[0])
	y2, _ := strconv.Atoi(p2[1])
	
	return x1, y1, x2, y2
}