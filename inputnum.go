package main

import "fmt"

func main() {
	var num1 int
	var num2  int

	fmt.Printf("Please Enter a large number : ")

	fmt.Scan(&num1)

	fmt.Printf("Please enter a smaller number : ")

	fmt.Scan(&num2)

	// % is the modul

	fmt.Println("the remainder of ", num1, " divide by ", num2, " is ", num1%num2)
}