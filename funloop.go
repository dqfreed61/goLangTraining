

package main

import "fmt"

func main() {

	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println(i, " -- FIZZ")
		}else  if i%5 ==0 {
			fmt.Println(i, " -- BUZZ")
		} else if i%15 ==0 {
			fmt.Println(i, " -- FIZZBUZZ")
		} else {
			fmt.Println(i)
		}


	}
	}


