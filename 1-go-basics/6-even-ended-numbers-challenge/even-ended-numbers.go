package main

import "fmt"

/*
	An even ended number is a number where the first and last digits are the same.

	Count how many even ended numbers are there that are a multiplication of two
	4 digit numbers.
*/

func main() {
	count := 0

	for a := 1000; a <= 9999; a++ {
		for b := a; b <= 9999; b++ {
			n := a * b
			s := fmt.Sprintf("%d", n)
			if s[0] == s[len(s)-1] {
				count++
			}
		}
	}

	fmt.Println(count)
}
