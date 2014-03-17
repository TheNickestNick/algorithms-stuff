package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func main() {
	num := big.NewInt(0)
	num.Exp(big.NewInt(2), big.NewInt(1000), nil)
	
	str := num.String()
	
	sum := 0
	for _, c := range str {		
		i, _ := strconv.Atoi(string(c))
		sum += i
	}
	
	fmt.Println(sum)
}