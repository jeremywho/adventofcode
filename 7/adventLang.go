package adventLang

import (
	"strconv"
	"strings"
)

// AdventProgram lets you execute instructions on it
type AdventProgram struct {
	memory map[string]chan uint16
	lines  []string
}

// NewAdventProgram allocates a new program.
func NewAdventProgram(lines []string) *AdventProgram {
	p := new(AdventProgram)
	p.memory = make(map[string]chan uint16)
	p.lines = lines

	for _, line := range p.lines {
		if len(line) > 0 {
			parts := strings.Split(line, " -> ")
			varName := parts[len(parts)-1]
			p.memory[varName] = make(chan uint16)
		}
	}

	return p
}

//Run will execute the statements passed in with NewAdventProgram
func (p *AdventProgram) Run() {
	for _, line := range p.lines {
		p.processStatement(line)
	}
}

//GetMemory returns the map of the memory for the program
func (p *AdventProgram) GetMemory() map[string]chan uint16 {
	return p.memory
}

//ProcessStatement will process a line of code
//once we get the value for a given variable (channel)
//we keep looping so we can feed that value to any
//other processes that need the value
//
//note: we could add a channel to close the go routines when they
//are no longer needed.
func (p *AdventProgram) processStatement(s string) {
	destToken := strings.Split(s, "->")

	dest := strings.TrimSpace(destToken[len(destToken)-1])
	tokens := strings.Split(strings.TrimSpace(destToken[0]), " ")

	//num assignment
	numTokens := len(tokens)
	if numTokens == 1 {
		l := strings.TrimSpace(tokens[0])
		if n, err := strconv.Atoi(l); err == nil {
			go func() {
				for {
					p.memory[dest] <- uint16(n)
				}
			}()
		} else {
			if memVal, ok := p.memory[l]; ok {
				go func() {
					chanVal := <-memVal
					for {
						p.memory[dest] <- chanVal
					}
				}()
			}
		}
	}

	//handle NOT
	//NOT x -> y
	if numTokens == 2 && (tokens[0] == "NOT") {
		go func() {
			r := strings.TrimSpace(tokens[1])

			//parse
			var x uint16

			if xi, err := strconv.Atoi(r); err == nil {
				//convert number string to number
				x = uint16(xi)
			} else {
				//get existing value
				if memVal, ok := p.memory[r]; ok {
					x = <-memVal
				}
			}

			for {
				p.memory[dest] <- ^x
			}
		}()
	}

	//x AND y -> z
	//x OR y -> z
	//x LSHIFT 2 -> z
	//x RSHIFT 2 -> z
	if numTokens == 3 {
		go func() {

			l := strings.TrimSpace(tokens[0])
			r := strings.TrimSpace(tokens[2])

			var x, y uint16

			//parse left
			if xi, err := strconv.Atoi(l); err == nil {
				//parsed number string
				x = uint16(xi)
			} else {
				//get existing value
				if memVal, ok := p.memory[l]; ok {
					x = <-memVal
				}
			}

			//parse right
			if yi, err := strconv.Atoi(r); err == nil {
				//parsed number string
				y = uint16(yi)
			} else {
				if memVal, ok := p.memory[r]; ok {
					//get existing value
					y = <-memVal
				}
			}

			//store result of bitwise AND
			if tokens[1] == "AND" {
				for {
					p.memory[dest] <- x & y
				}
			}

			//store result of bitwise OR
			if tokens[1] == "OR" {
				for {
					p.memory[dest] <- x | y
				}
			}

			//store result of bitwise LSHIFT
			if tokens[1] == "LSHIFT" {
				for {
					p.memory[dest] <- x << y
				}
			}

			//store result of bitwise RSHIFT
			if tokens[1] == "RSHIFT" {
				for {
					p.memory[dest] <- x >> y
				}
			}
		}()
	}
}
