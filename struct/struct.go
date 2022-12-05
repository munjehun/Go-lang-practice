package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	p := person{}
	p.name = "Lee"
	p.age = 10

	p2 := person{name: "Seam", age: 49}

	p3 := new(person) //new를 사용하면 모든 필드를 0으로 초기화
	p3.name = "Lee"

	fmt.Println(p)
	fmt.Println(p2)
	fmt.Println(p3)
}
