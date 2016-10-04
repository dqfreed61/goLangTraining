package main

import "fmt"

func main() {
	var fname string
	var lname string
	fmt.Printf("Enter your first and last name")

	fmt.Scan(&fname)
	fmt.Scan(&lname)

	fmt.Println("Hello", fname, " ", lname)
}