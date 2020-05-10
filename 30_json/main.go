//Go offers built-in support for JSON encoding and decoding,
//including to and from built-in and custom data types.

package main

import (
	"encoding/json"
	"fmt"
)

type student struct {
	//Only exported fields will be encoded/decoded in JSON.
	//Fields must start with capital letters to be exported.
	Name   string
	RollNo int `json:"RegNo"` //Json msg will contain RegNo, instead of RollNo
}

func main() {
	//Marshal will convert data to json msg
	fmt.Println("Marshalling examples")
	str, _ := json.Marshal("Golang")
	fmt.Println(str)         //generates []byte containing json msg
	fmt.Println(string(str)) //convert to string

	intgr, _ := json.Marshal(1)
	fmt.Println(string(intgr))

	slice := []string{"hi", "welcome", "to", "go", "world"}
	sliceB, _ := json.Marshal(slice)
	fmt.Println(string(sliceB))

	maps := map[string]int{"One": 1, "Two": 2, "Three": 3}
	mapB, _ := json.Marshal(maps)
	fmt.Println(string(mapB))

	stud := &student{"Jon", 1}
	studentB, _ := json.Marshal(stud)
	fmt.Println(string(studentB))

	fmt.Println()
	//Unmarshal will convert json msg to data
	//Above marshled example will be unmarshaled here
	fmt.Println("Unmarshalling examples")
	stud1 := student{}
	json.Unmarshal(studentB, &stud1)
	fmt.Println(stud1)

	map1 := map[string]int{}
	json.Unmarshal(mapB, &map1)
	fmt.Println(map1)

	slice1 := []string{}
	json.Unmarshal(sliceB, &slice1)
	fmt.Println(slice1)
}
