package main

import "fmt"

func ReturnId() (id int, err error) {
	defer func(id int) {
		fmt.Println(id)
		if id == 10 {
			err = fmt.Errorf("Invalid Id\n")
		}
	}(id)

	id = 10

	return
}

func main() {
	id, err := ReturnId()
	fmt.Println(id, err)
}
