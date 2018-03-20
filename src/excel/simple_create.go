package main

import (
	"github.com/360EntSecGroup-Skylar/excelize"
	"fmt"
)

func main() {
	xlsx := excelize.NewFile()
	index := xlsx.NewSheet("Sheet2")

	xlsx.SetCellValue("Sheet2", "A2", "Hello, world.")
	xlsx.SetCellValue("Sheet1", "B2", 100)
	xlsx.SetActiveSheet(index)
	err := xlsx.SaveAs("F:\\github\\go_learning\\src\\excel\\new\\simple_create.xlsx")
	if err != nil {
		fmt.Println(err)
	    return
	}
}
