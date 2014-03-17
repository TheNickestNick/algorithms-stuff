package main

import (
	"fmt"
)

func factor(n int) map[int]int {
	factors := make(map[int]int)
	d := 2
	
	for n > 1 {
		for n % d == 0 {
			_, present := factors[d]
			if !present {
				factors[d] = 0
			}
			
			factors[d] += 1
			n /= d
		}
		
		d++
		
		if d * d > n && n > 1 {
			if _, present := factors[d]; !present {
				factors[n] = 0
			}
			factors[n] += 1
			break
		}
	}
	
	return factors
}

func triangular(n int) int {
	return (n * (n + 1)) / 2
}

func countDivisors(n int) int {
	count := 1
	for _, power := range factor(n) {
		count *= (power + 1)
	}
	return count
}

func main() {
	for n := 1; ; n++ {
		tri := triangular(n)
		count := countDivisors(tri)
		
		if (count > 500) {
			fmt.Println("found it:", tri)
			fmt.Println(count, "divisors")
			return
		}
	}
}