package main

import (
	"encoding/json"
	"log"
	"fmt"
)

func main() {
	var i *int = nil
	bytes, err := json.Marshal(i)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println(string(bytes))

	var k int = 1
	var j *int =  &k
	json.Unmarshal(bytes, j)
	fmt.Println(*j)
}
