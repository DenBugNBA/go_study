package main

import (
	"encoding/json"
	"fmt"
)

type someStruct struct {
	Name   string
	Age    int
	Status bool
	Values []int
}

func main() {
	s := someStruct{
		Name:   "John Connor",
		Age:    35,
		Status: true,
		Values: []int{15, 11, 37},
	}

	data, err := json.MarshalIndent(s, "", "\t")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s", data)

	//{
	//	"Name": "John Connor",
	//	"Age": 35,
	//	"Status": true,
	//	"Values": [
	//		15,
	//		11,
	//		37
	//	]
	//}
}
