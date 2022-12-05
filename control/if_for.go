package main

import "fmt"

func printhello(x int) int {

	count := 0
	for i := 0; i < x; i++ {
		if i%2 == 0 {
			fmt.Println("Hello bstudent")
			count++
		}
	}
	return count
}

func main() {
	fmt.Println("count :", printhello(10))
}
