package main

import (
	"github.com/zoumo/goset"
	"fmt"
)

func main() {
	ints := []int{1, 2, 3, 4}
	floats := []float64{1.0, 2.0, 3.0}
	strings := []string{"1", "2", "3"}
	set1 := goset.NewSetFrom(ints)
	set2 := goset.NewSafeSetFrom(ints)
	set3 := goset.NewSetFrom(floats)
	set4 := goset.NewSafeSetFrom(floats)
	set5 := goset.NewSetFrom(strings)
	set6 := goset.NewSafeSetFrom(strings)

	fmt.Println(set1, set2, set3, set4, set5, set6)

	set7 := goset.NewSet()
	set7.Add(1)
	set7.Add(2)
	fmt.Println(set7)
	fmt.Println(set7.Elements())
}
