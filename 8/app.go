package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

// using quote/unquote is a better soln
// from reddit.com/u/Astrus
func quote(str string) string {
	return strconv.Quote(str)
}

func unquote(str string) string {
	s, _ := strconv.Unquote(str)
	return s
}

func main() {
	b, _ := ioutil.ReadFile("8\\input.txt")
	input := string(b[:])
	lines := strings.Split(input, "\n")

	total, values, encoded := 0, 0, 0

	for _, line := range lines {
		//had to add this check due to empty line in input file
		if len(line) < 1 {
			continue
		}

		total += len(line)
		encoded += len(quote(line))

		result := strings.Replace(line, `\\`, "x", -1)
		result = strings.Replace(result, `\"`, "x", -1)
		result = strings.Replace(result, `"`, "", -1)
		r, _ := regexp.Compile(`\\x[a-fA-F0-9][a-fA-F0-9]`)
		result = r.ReplaceAllString(result, "x")

		values += len(result)
	}

	fmt.Println(total, " - ", values, " = ", total-values)
	fmt.Println(encoded, " - ", total, " = ", encoded-total)
}
