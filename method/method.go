package main

import "fmt"

type Rect struct { //struct 정의
	width, height int
}

func (r Rect) area() int { //Rect의 area() 메소드
	r.width++
	return r.width * r.height
}
func (r *Rect) area2() int { //포인터 Receiver
	r.width++
	return r.width * r.height

}
func main() {
	rect := Rect{10, 20}
	area1 := rect.area()           //메서드 호출
	fmt.Println(rect.width, area1) //10, 220 => 데이터 안변함
	area2 := rect.area2()          //메서드 호출
	fmt.Println(rect.width, area2) //11, 220 => 데이터 1증가
}
