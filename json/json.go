package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Task struct {
	Title  string `json:"title"`
	Status status
}

type status int

const (
	UNKNOWN status = iota
	TODO
	DONE
)

func main() {
	ExampleTask_marshalJSON()
	ExampleTask_unmarshalJSON()
}

func ExampleTask_marshalJSON() {
	t := Task{
		"Laundry",
		DONE,
	}
	b, err := json.Marshal(t)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(b)) //{"title":"Laundry","Status":2}
}

func ExampleTask_unmarshalJSON() {
	b := []byte(`{"Title":"Buy Milk","Status":2} `)
	t := Task{} //빈 구조체
	err := json.Unmarshal(b, &t)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(t.Title)  //Buy Milk
	fmt.Println(t.Status) //2
}
