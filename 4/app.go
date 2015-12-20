package main

import (
	"crypto/md5"
	"fmt"
	"strings"
)

func main() {
	secret := "bgvyzdsv"	
	
	for i := 100000; i <= 9999999; i++ {
		t := fmt.Sprintf("%s%d", secret, i)
		d := []byte(t)
		hash := md5.Sum(d)
		ps := string(hash[:])
		hs := fmt.Sprintf("%x\n", ps)
		
		if strings.HasPrefix(hs, "00000") {
			fmt.Printf("%d\n", i)
			fmt.Printf("%x", hash)			
		}
		
		if strings.HasPrefix(hs, "000000") {
			fmt.Printf("\n\nPart 2\n")
			fmt.Printf("%d\n", i)
			fmt.Printf("%x", hash)	
			return		
		}
	}
}