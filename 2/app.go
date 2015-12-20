package main

import (
	"fmt"
	"strings"
	"io/ioutil"
	"strconv"
	"sort"
	"sync"
)

func main() {
	//var total int
	paperChan := make(chan int)
	ribbonChan := make(chan int)
	readTotalPaper := make(chan int)
	readTotalRibbon := make(chan int)
	
	var wg sync.WaitGroup
	
	go func (paper chan int, readPaper chan int, ribbon chan int, readRibbon chan int) {
		totalPaper := 0
		totalRibbon := 0
		for {
			select {
            	case p := <- paper:
                	totalPaper += p
				case r := <- ribbon:
					totalRibbon += r
				case readPaper <- totalPaper:
				case readRibbon <- totalRibbon:
        	}
		}
	}(paperChan, readTotalPaper, ribbonChan, readTotalRibbon)
	
	b, _ := ioutil.ReadFile("input.txt")
	input := string(b[:])
	ss := strings.Split(input, "\n")
	
	wg.Add(len(ss))
	
	for _, box := range ss {
		go processBox(box, paperChan, ribbonChan, &wg)
	}
	
	wg.Wait()	
	fmt.Println("Done")
	fmt.Printf("Total Paper: %d\n", <-readTotalPaper)
	fmt.Printf("Total Ribbon: %d\n", <-readTotalRibbon)
}

func processBox(b string, paperChan chan int, ribbonChan chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	
	if(len(b) < 1) {
		return
	}
	
	vals := strings.Split(b, "x")
	l, _ := strconv.Atoi(vals[0])
	w, _ := strconv.Atoi(vals[1])
	h, _ := strconv.Atoi(vals[2])
	
	s := []int{l, w, h}
	sort.Ints(s)
	//fmt.Println(b)
	paperChan <- (2*l*w) + (2*w*h) + (2*h*l) + (s[0]*s[1])
	ribbonChan <- (l*w*h) + (2*s[0]) + (2*s[1])
}