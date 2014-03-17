package main

import (
	"fmt"
)

func collatzLength(n int) int {
	count := 1
	
	for n > 1 {
		if n % 2 == 0 {
			n /= 2
		} else {
			n = (3 * n) + 1
		}
		
		count++
	}
	
	return count
}

func main() {
	max, maxlen := 0, 0
	
	for n := 2; n < 1000000; n++ {
		len := collatzLength(n)
		
		if len > maxlen {
			max, maxlen = n, len
		}
	}

	fmt.Println("number with largest collatz length:", max)
	fmt.Println("collatz length:", maxlen)
}