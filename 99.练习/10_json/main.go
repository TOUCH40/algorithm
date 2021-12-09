package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name      string `json:"name"`
	BirthDate string `json:"birth_date"`
}

func main() {
	jsonPerson := `{"name":"小明","birth_date":"1990-08-03 00:00:00"}`
	var p Person
	json.Unmarshal([]byte(jsonPerson), &p)
	fmt.Printf("%v\n", p)
	js, err := json.Marshal(p)
	if err == nil {
		fmt.Printf("%v\n", string(js))
	}
}
