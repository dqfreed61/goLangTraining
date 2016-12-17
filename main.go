package main

import "fmt"


// if we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of
// these multiples is 23
// find the sum of all the multiples of 3 or 5 below 1000

func threeorfive (value int) int {

	var sums int
	for i := 1; i < value; i++ {
		if (i%3 ==0 || i%5 ==0) {
			sums = sums + i
		//	fmt.Println(i)
		}
	}
	return sums
}



func main() {

	total := threeorfive(1000)
	fmt.Println(total)
}