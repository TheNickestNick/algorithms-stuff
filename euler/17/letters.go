package main

import "fmt"
import "strings"

var lessThan20 map[int]string
var tensDigits map[int] string

func init() {
	lessThan20 = map[int]string {
		0: "zero", 1: "one", 2: "two", 3: "three", 4: "four", 
		5: "five", 6: "six", 7: "seven", 8: "eight", 9: "nine", 
		10: "ten", 11: "eleven", 12: "twelve", 13: "thirteen",
		14: "fourteen", 15: "fifteen", 16: "sixteen",
		17: "seventeen", 18: "eighteen", 19: "nineteen",
	}

	tensDigits = map[int]string {
		2: "twenty", 3: "thirty", 4: "forty", 5: "fifty",
		6: "sixty", 7: "seventy", 8: "eighty", 9: "ninety",
	}
}

func intToString(n int) string {
	if n < 20 {
		return lessThan20[n]
	}

	if n < 100 {
		tens := n/10
		ones := n % 10
		str := tensDigits[tens]

		if ones > 0 {
			str += " " + lessThan20[ones]
		}
		
		return str
	}
	
	if (n < 1000) {
		hundreds := n / 100
		str := lessThan20[hundreds] + " hundred"
		
		if n % 100 > 0 {
			str += " and " + intToString(n % 100)
		}
		
		return str
	}
	
	if n == 1000 {
		return "one thousand"
	}

	return "INVALID"
}

func main() {
	str := ""
	for n := 1; n <= 1000; n++ {
		str += intToString(n)
		//fmt.Println(n, len(strings.Replace(str, " ", "", -1)), str)
	}
	
	fmt.Println(len(strings.Replace(str, " ", "", -1)))
}