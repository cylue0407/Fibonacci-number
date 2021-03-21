/*
* Implement Fibonacci number using Golang Recursive
 */
package main

import (
	"fmt"
)

func main() {
	usrInput := 0
	fmt.Println("Please input Fibonacci number you want: ")
	fmt.Scanln(&usrInput)
	fmt.Println("Result:", GetFibNum(usrInput))
}

func GetFibNum(n int) int {
	if n > 0 && n < 3 {
		if n == 1 {
			return 1
		} else {
			return 2
		}
	} else if n > 2 {
		return GetFibNum(n-1) + GetFibNum(n-2)
	} else {
		return 0
	}
}
