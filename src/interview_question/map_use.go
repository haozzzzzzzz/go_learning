package main

import "fmt"

type student struct {
	Name string
	Age int
}


func pase_student() {
	m := make(map[string]*student)
	stus := []student {
		{
			Name: "zhou",
			Age: 24,
		},
		{
			Name: "run",
			Age: 23,
		},
		{
			Name: "fa",
			Age: 22,
		},
	}
	for _, stu := range stus {
		m[stu.Name] = &stu // 这里浅复制
	}
	fmt.Println(m)
}

func main()  {
	pase_student()
}