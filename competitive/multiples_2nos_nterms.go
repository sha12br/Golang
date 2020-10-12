
// If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.

// Finish the solution so that it returns the sum of all the multiples of 3 or 5 below the number passed in. 
// Note: If the number is a multiple of both 3 and 5, only count it once. Also, if a number is negative, return 0(for languages that do have them)


package main

import (
	"fmt"
	"math"
)

func multi(n float64, d float64) int {
	nn := int(math.Floor(n/d))	
	return (nn*((2*int(d)+ (nn-1)*int(d))))/2
}

func calc_multi(n float64, n1 float64, n2 float64) int {

	if n < 0|| n1<0 || n2 < 0 {
		return 0
	} else {
		return (multi(n,n1)+multi(n,n2)-multi(n,n1*n2))
	}
}

func main() {
	fmt.Println(calc_multi(15,-3,5))
}
