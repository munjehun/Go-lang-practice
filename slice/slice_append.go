package main

import "fmt"

func main() {
	f1 := []string{"사과", "바나나", "토마토"}
	f2 := []string{"포도", "딸기"}
	f3 := append(f1, f2...)
	f4 := append(f1[:2], f2...)

	fmt.Println(f1) //[사과 바나나 토마토]
	fmt.Println(f2) //[포도 딸기]
	fmt.Println(f3) //[사과 바나나 토마토 포도 딸기]
	fmt.Println(f4) //[사과 바나나 포도 딸기]
}
