package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Person : struct for a person
type Person struct {
	First string
	Last  string
	Age   int
}

func main() {
	var p1 Person

	rdr := strings.NewReader(`{"First":"James", "Last":"Bond", "Age":20}`)
	json.NewDecoder(rdr).Decode(&p1)

	fmt.Println("---------------")
	fmt.Println(p1.First)
	fmt.Println(p1.Last)
	fmt.Println(p1.Age)
	fmt.Printf("%T \n", p1)
}
