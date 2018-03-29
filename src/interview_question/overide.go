package main

import "fmt"

type People struct {
}

func (m *People) ShowA() {
	fmt.Println("show A")
	m.ShowB()
}

func (m *People) ShowB()  {
	fmt.Println("show B")
}

type Teacher struct {
	People
}

func (m *Teacher) ShowB()  {
	fmt.Println("teacher showB")
}

func main()  {
	t := &Teacher{}
	t.ShowA()
	t.ShowB()
}

/**
show A
show B
teacher showB
 */